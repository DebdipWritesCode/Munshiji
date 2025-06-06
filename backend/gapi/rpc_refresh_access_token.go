package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) RefreshAccessToken(ctx context.Context, _ *emptypb.Empty) (*pb.LoginUserResponse, error) {
	refreshMtdt := server.extractRefreshMetadata(ctx)
	if refreshMtdt.token == "" {
		return nil, status.Error(codes.Unauthenticated, "missing refresh token in metadata")
	}

	payload, err := server.tokenMaker.VerifyToken(refreshMtdt.token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token: %s", err.Error())
	}

	session, err := server.store.GetSessionsByRefreshToken(ctx, refreshMtdt.token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get session: %s", err.Error())
	}

	if session.RefreshToken != refreshMtdt.token || session.UserID != payload.UserID {
		return nil, status.Error(codes.Unauthenticated, "refresh token does not match session")
	}

	if time.Now().After(session.RefreshExpiresAt) {
		return nil, status.Error(codes.Unauthenticated, "refresh token has expired")
	}

	accessTokenDurationStr := server.config.AccessTokenDuration // 15m

	accessTokenDuration, err := time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid access token duration: %s", err.Error())
	}

	newAccessToken, err := server.tokenMaker.CreateToken(payload.UserID, accessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err.Error())
	}

	arg := db.UpdateSessionExpirationParams{
		ID:        session.ID,
		ExpiresAt: sql.NullTime{Time: time.Now().Add(accessTokenDuration), Valid: true},
		JwtToken:  newAccessToken,
	}

	updatedSession, err := server.store.UpdateSessionExpiration(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update session expiration: %s", err.Error())
	}

	user, err := server.store.GetUserByID(ctx, updatedSession.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user with ID %d not found", updatedSession.UserID)
		}
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err.Error())
	}

	mtdt := server.extractMetadata(ctx)
	rsp := &pb.LoginUserResponse{
		User:     converUserToProto(user),
		Metadata: convertMetadataToProto(mtdt),
		JwtToken: newAccessToken,
	}

	return rsp, nil
}
