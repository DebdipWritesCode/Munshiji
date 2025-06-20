package gapi

import (
	"fmt"

	db "github.com/DebdipWritesCode/Munshiji/backend/db/sqlc"
	"github.com/DebdipWritesCode/Munshiji/backend/pb"
	"github.com/DebdipWritesCode/Munshiji/backend/token"
	"github.com/DebdipWritesCode/Munshiji/backend/util"
	"github.com/DebdipWritesCode/Munshiji/backend/worker"
)

type Server struct {
	pb.UnimplementedMunshijiServer
	config          util.Config
	tokenMaker      token.Maker
	store           db.Store
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
