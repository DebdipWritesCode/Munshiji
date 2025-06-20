package gapi

import (
	db "github.com/DebdipWritesCode/Munshiji/backend/db/sqlc"
	"github.com/DebdipWritesCode/Munshiji/backend/pb"
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

func convertScoreSheetToProto(sheet db.ScoreSheet) *pb.ScoreSheet {
	return &pb.ScoreSheet{
		Id:            sheet.ID,
		Name:          sheet.Name,
		CommitteeName: sheet.CommitteeName,
		Chair:         sheet.Chair,
		ViceChair:     sheet.ViceChair.String,
		Rapporteur:    sheet.Rapporteur.String,
		CreatedBy:     sheet.CreatedBy,
		CreatedAt: func() *timestamppb.Timestamp {
			if sheet.CreatedAt.Valid {
				return timestamppb.New(sheet.CreatedAt.Time)
			}
			return nil
		}(),
		UpdatedAt: func() *timestamppb.Timestamp {
			if sheet.UpdatedAt.Valid {
				return timestamppb.New(sheet.UpdatedAt.Time)
			}
			return nil
		}(),
	}
}

func ConvertDelegateToProto(delegate db.Delegate) *pb.Delegate {
	return &pb.Delegate{
		Id:           delegate.ID,
		ScoreSheetId: delegate.ScoreSheetID,
		Name:         delegate.Name,
	}
}

func ConvertScoreToProto(score db.Score) *pb.Score {
	return &pb.Score{
		Id:          score.ID,
		DelegateId:  score.DelegateID,
		ParameterId: score.ParameterID,
		Value:       score.Value,
		Note:        score.Note.String,
	}
}

func ConvertParameterToProto(parameter db.Parameter) *pb.Parameter {
	return &pb.Parameter{
		Id:                 parameter.ID,
		ScoreSheetId:       parameter.ScoreSheetID,
		Name:               parameter.Name,
		RuleType:           parameter.RuleType,
		IsSpecialParameter: parameter.IsSpecialParameter,
		SpecialScoresRule:  parameter.SpecialScoresRule.String,
		SpecialLengthRule:  parameter.SpecialLengthRule.String,
		ScoreWeight:        parameter.ScoreWeight.Float64,
		LengthWeight:       parameter.LengthWeight.Float64,
	}
}
