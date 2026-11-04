package httputil

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func SendJSON(w http.ResponseWriter, code int, message proto.Message) {
	payload, err := (protojson.MarshalOptions{
		UseProtoNames: true,
	}).Marshal(message)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(payload)
}

func SendGRPCError(w http.ResponseWriter, err error) {
	switch status.Code(err) {
	case codes.InvalidArgument:
		http.Error(w, status.Convert(err).Message(), http.StatusBadRequest)
	case codes.AlreadyExists, codes.Aborted, codes.FailedPrecondition:
		http.Error(w, "operation conflicts with current state", http.StatusConflict)
	case codes.Unavailable:
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
	case codes.DeadlineExceeded:
		http.Error(w, "service timeout", http.StatusGatewayTimeout)
	default:
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
