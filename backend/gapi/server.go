package gapi

import (
	"fmt"

	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/token"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/util"
)

type Server struct {
	pb.UnimplementedMunshijiServer
	config     util.Config
	tokenMaker token.Maker
	store      db.Store
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
