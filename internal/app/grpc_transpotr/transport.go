package grpc_transpotr

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

const (
	grpcHost = "localhost"
	grpcPort = 50051
)

// TransportService транспортный слой.
type TransportService struct {
	net.Listener
}

// NewTransportService возвращается транспортный слой.
func NewTransportService() (*TransportService, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", grpcHost, grpcPort))
	if err != nil {
		return nil, errors.Wrap(err, "failed to listen")
	}

	return &TransportService{listener}, nil
}
