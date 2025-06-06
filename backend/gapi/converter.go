package gapi

import (
	db "github.com/DebdipWritesCode/MUN_Scoresheet/backend/db/sqlc"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func converUserToProto(user db.User) *pb.User {
	return &pb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		CreatedAt: func() *timestamppb.Timestamp {
			if user.CreatedAt.Valid {
				return timestamppb.New(user.CreatedAt.Time)
			}
			return nil
		}(),
	}
}

func convertMetadataToProto(m *Metadata) *pb.MetaData {
	if m == nil {
		return nil
	}
	return &pb.MetaData{
		UserAgent: m.UserAgent,
		ClientIp:  m.ClientIP,
	}
}
