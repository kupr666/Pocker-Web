package lobby

import (
	"encoding/json"
	"net/http"

	tablesv1 "github.com/kupr666/Pocker-Web/gen/go/tables/v1"
	"github.com/kupr666/Pocker-Web/services/gateway/internal/interfaces/http/httputil"
)

type GatewayServer struct {
	lobby tablesv1.LobbyServiceClient
}

func NewGatewayServer(client tablesv1.LobbyServiceClient) *GatewayServer {
	return &GatewayServer{lobby: client}
}

func (g *GatewayServer) CreateTables(w http.ResponseWriter, r *http.Request) {
	var req createTableRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	// validation with validator lib

	// call lobby client through grpc
	response, err := g.lobby.CreateTable(r.Context(), &tablesv1.CreateTableRequest{})
	if err != nil {
		httputil.SendGRPCError(w, err)
		return
	}
	httputil.SendJSON(w, http.StatusCreated, response)
}
