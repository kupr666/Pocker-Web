package main

import (
	"context"
	"strings"

	"github.com/google/uuid"
	gamev1 "github.com/kupr666/Pocker-Web/gen/go/game/v1"
	v1 "github.com/kupr666/Pocker-Web/gen/go/tables/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LobbyServer struct {
	v1.UnimplementedLobbyServiceServer
	game gamev1.GameServiceClient
}

func (s *LobbyServer) CreateTable(
	ctx context.Context,
	req *v1.CreateTableRequest,
) (*v1.CreateTableResponse, error) {

	if err := validateCreateTableRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tableID := uuid.NewString()
	rules, err := prepareRules(req.Rules)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mode := req.GetAccessMode()
	if mode == v1.AccessMode_ACCESS_MODE_UNSPECIFIED {
		mode = v1.AccessMode_ACCESS_MODE_PUBLIC
	}
	createdAt := timestamppb.Now()

	_, err = s.game.CreateTableRuntime(ctx, &gamev1.CreateTableRuntimeRequest{
		TableId: tableID,
		Rules:   rules,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateTableResponse{
		Table: &v1.Table{
			TableId:       tableID,
			Name:          strings.TrimSpace(req.GetName()),
			OwnerUserId:   req.GetOwnerUserId(),
			AccessMode:    mode,
			Rules:         rules,
			OccupiedSeats: 0,
			CreatedAt:     createdAt,
		},
	}, nil
}
