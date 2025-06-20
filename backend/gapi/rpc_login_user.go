package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/DebdipWritesCode/Munshiji/backend/db/sqlc"
	"github.com/DebdipWritesCode/Munshiji/backend/pb"
	"github.com/DebdipWritesCode/Munshiji/backend/util"
	"github.com/DebdipWritesCode/Munshiji/backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginUserRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user with email %s not found %s", req.GetEmail(), err.Error())
		}
	}

	if user.IsEmailVerified != true {
		return nil, status.Errorf(codes.FailedPrecondition, "email for user with email %s is not verified", req.GetEmail())
	}

	passwordToCheck := user.PasswordHash
	if err := util.CheckPassword(passwordToCheck, req.GetPassword()); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password for user with email %s: %s", req.GetEmail(), err.Error())
	}

	accessTokenDurationStr := server.config.AccessTokenDuration   // 15m
	refreshTokenDurationStr := server.config.RefreshTokenDuration // 48h

	accessTokenDuration, err := time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid access token duration: %s", err.Error())
	}

	refreshTokenDuration, err := time.ParseDuration(refreshTokenDurationStr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid refresh token duration: %s", err.Error())
	}

	accessToken, err := server.tokenMaker.CreateToken(user.ID, accessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err.Error())
	}

	refreshToken, err := server.tokenMaker.CreateToken(user.ID, refreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token: %s", err.Error())
	}

	arg := db.CreateSessionParams{
		UserID:           user.ID,
		JwtToken:         accessToken,
		RefreshToken:     refreshToken,
		ExpiresAt:        sql.NullTime{Time: time.Now().Add(accessTokenDuration), Valid: true},
		RefreshExpiresAt: time.Now().Add(refreshTokenDuration),
	}

	session, err := server.store.CreateSession(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session: %s", err.Error())
	}

	if err := setRefreshTokenCookie(ctx, session.RefreshToken, refreshTokenDuration); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set refresh token cookie: %s", err.Error())
	}

	mtdt := server.extractMetadata(ctx)
	rsp := &pb.LoginUserResponse{
		User:     converUserToProto(user),
		Metadata: convertMetadataToProto(mtdt),
		JwtToken: accessToken,
	}

	return rsp, nil
}

func validateLoginUserRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	return violations
}

func setRefreshTokenCookie(ctx context.Context, token string, duration time.Duration) error {
	cookie := fmt.Sprintf("refresh_token=%s; HttpOnly; Path=/; Max-Age=%d; Secure; SameSite=Strict", token, int(duration.Seconds()))
	header := metadata.Pairs("Set-Cookie", cookie)

	err := grpc.SendHeader(ctx, header)
	if err != nil {
		return err
	}

	return nil
}
