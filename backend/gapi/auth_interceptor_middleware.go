package gapi

import (
	"context"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/token"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeaderKey = "authorization"
	bearerPrefix           = "bearer"
)

var publicMethods = map[string]bool{
	"/pb.Munshiji/LoginUser":        true,
	"/pb.Munshiji/CreateUser":       true,
	"/pb.Munshiji/RenewAccessToken": true,
}

func AuthFunc(tokenMaker token.Maker) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		stream := grpc.ServerTransportStreamFromContext(ctx)
		if stream != nil {
			fullMethod := stream.Method()
			if publicMethods[fullMethod] {
				return ctx, nil
			}
		}

		accessToken, err := grpc_auth.AuthFromMD(ctx, bearerPrefix)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to extract token from metadata: %v", err)
		}

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		newCtx := context.WithValue(ctx, "tokenPayload", payload)
		return newCtx, nil
	}
}
