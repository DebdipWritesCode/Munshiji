package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	_ "github.com/DebdipWritesCode/MUN_Scoresheet/backend/doc/statik"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/gapi"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/token"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/util"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create token maker")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	store := db.NewStore(conn)

	go runGatewayServer(store, config)
	runGrpcServer(store, config, tokenMaker)
}

func runGrpcServer(store db.Store, config util.Config, tokenMaker token.Maker) {
	server, err := gapi.NewServer(config, store)
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

	log.Info().Msgf("gRPC server listening on %s", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start gRPC server")
	}
}

func runGatewayServer(store db.Store, config util.Config) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create gRPC server")
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen on %s", config.HTTPServerAddress)
	}

	log.Info().Msgf("HTTP Gateway server listening on %s", config.HTTPServerAddress)
	handler := gapi.HttpLogger(corsHandler)

	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start HTTP server")
	}
}
