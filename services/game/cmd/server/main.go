package main

import (
	"context"
	"log"
	"net"
	"sync"

	v1 "github.com/kupr666/Pocker-Web/gen/go/game/v1"
	tablev1 "github.com/kupr666/Pocker-Web/gen/go/table/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type TableRuntime struct {
	Rules *tablev1.TableRules
	Seats map[uint32]string
}

type GameServer struct {
	v1.UnimplementedGameServiceServer
	mu     sync.Mutex
	tables map[string]*TableRuntime // table_id - > table condition
}

func (s *GameServer) CreateTableRuntime(
	ctx context.Context,
	req *v1.CreateTableRuntimeRequest,
) (*v1.CreateTableRuntimeResponse, error) {
	id := req.GetTableId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "table_id is required")
	}
	if req.GetRules() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "rules are required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if existing, found := s.tables[id]; found {
		if !proto.Equal(existing.Rules, req.GetRules()) {
			return nil, status.Errorf(codes.AlreadyExists, "table_id already exists with different rules")
		}
		return &v1.CreateTableRuntimeResponse{TableId: id}, nil
	}
	s.tables[id] = &TableRuntime{
		Rules: proto.Clone(req.GetRules()).(*tablev1.TableRules),
		Seats: make(map[uint32]string),
	}
	return &v1.CreateTableRuntimeResponse{
		TableId: req.GetTableId(),
	}, nil
}

func main() {
	// quick implementation to get minimal version of get table
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	// регистрирует в gRPC-сервере методы GameService, объявленные в .proto.
	// Она не регистрирует все произвольные методы структуры GameServer.
	v1.RegisterGameServiceServer(server, &GameServer{
		tables: make(map[string]*TableRuntime),
	})
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
