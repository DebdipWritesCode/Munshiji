package gapi

import (
	"context"
	"strings"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/token"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeaderKey = "authorization"
	bearerPrefix           = "bearer "
)

func authFunc(tokenMaker token.Maker) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		vals := md.Get(authorizationHeaderKey)
		if len(vals) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
		}

		authHeader := vals[0]
		if !strings.HasPrefix(strings.ToLower(authHeader), bearerPrefix) {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization format")
		}

		accessToken := strings.TrimPrefix(authHeader, bearerPrefix)
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		newCtx := context.WithValue(ctx, "tokenPayload", payload)
		return newCtx, nil
	}
}
