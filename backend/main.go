package main

import (
	"context"
	"database/sql"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	db "github.com/DebdipWritesCode/Munshiji/backend/db/sqlc"
	_ "github.com/DebdipWritesCode/Munshiji/backend/doc/statik"
	"github.com/DebdipWritesCode/Munshiji/backend/gapi"
	"github.com/DebdipWritesCode/Munshiji/backend/mail"
	"github.com/DebdipWritesCode/Munshiji/backend/pb"
	"github.com/DebdipWritesCode/Munshiji/backend/token"
	"github.com/DebdipWritesCode/Munshiji/backend/util"
	"github.com/DebdipWritesCode/Munshiji/backend/worker"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/rakyll/statik/fs"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create token maker")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	group, ctx := errgroup.WithContext(ctx)

	runTaskProcessor(ctx, config, group, redisOpt, store)
	runGatewayServer(ctx, store, config, group, taskDistributor)
	runGrpcServer(ctx, store, config, group, taskDistributor, tokenMaker)

	err = group.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error in goroutines:")
	}
}

func runTaskProcessor(ctx context.Context, config util.Config, group *errgroup.Group, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("Starting task processor...")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}

	group.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("Shutting down task processor...")

		taskProcessor.Shutdown()
		log.Info().Msg("Task processor shutdown complete")
		return nil
	})
}

func runGrpcServer(ctx context.Context, store db.Store, config util.Config, group *errgroup.Group, taskDistributor worker.TaskDistributor, tokenMaker token.Maker) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create gRPC server")
	}

	authInterceptor := gapi.AuthFunc(tokenMaker)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			gapi.GrpcLogger,
			auth.UnaryServerInterceptor(authInterceptor),
		),
	)
	pb.RegisterMunshijiServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen on %s", config.GRPCServerAddress)
	}

	group.Go(func() error {
		log.Info().Msgf("Starting GRPC server on %s...\n", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil // If the server is stopped gracefully, we return nil to indicate no error occurred.
			}
			log.Fatal().Err(err).Msg("cannot start gRPC server:")
			return err
		}

		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("Shutting down gRPC server...")
		grpcServer.GracefulStop() // This gracefully stops the gRPC server, allowing it to finish processing ongoing requests before shutting down.
		log.Info().Msg("gRPC server stopped.")
		return nil
	})
}

func runGatewayServer(ctx context.Context, store db.Store, config util.Config, group *errgroup.Group, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create gRPC server:")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	headerMatcher := runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		if key == "Authorization" || key == "Cookie" {
			return key, true
		}
		return runtime.DefaultHeaderMatcher(key)
	})

	cookieForwarder := runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
		if md, ok := runtime.ServerMetadataFromContext(ctx); ok {
			if cookies := md.HeaderMD.Get("Set-Cookie"); len(cookies) > 0 {
				for _, c := range cookies {
					w.Header().Add("Set-Cookie", c)
				}
			}
		}
		return nil
	})

	grpcMux := runtime.NewServeMux(jsonOption, headerMatcher, cookieForwarder)

	err = pb.RegisterMunshijiHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to register gRPC gateway server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFs, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create statik filesystem")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFs))
	mux.Handle("/swagger/", swaggerHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Cookie"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}).Handler(mux)

	httpServer := &http.Server{
		Handler: gapi.HttpLogger(corsHandler),
		Addr:    config.HTTPServerAddress,
	}

	group.Go(func() error {
		log.Info().Msgf("Starting HTTP Gateway server on %s...\n", httpServer.Addr)

		err := httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("cannot start HTTP Gateway server:")
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("Shutting down HTTP Gateway server...")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("cannot shutdown HTTP Gateway server:")
			return err
		}
		log.Info().Msg("HTTP Gateway server stopped.")
		return nil
	})
}
