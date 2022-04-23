package grpc

import (
	"context"
	"log"
	"net"

	minipulsa "github.com/elangreza14/minipulsa/authentication/adapter/grpc/minipulsa"
	port "github.com/elangreza14/minipulsa/authentication/port"

	grpc "google.golang.org/grpc"
)

// type Adapter interface {
// 	LoginRegister(context.Context, *LoginRegisterRequest) (*LoginRegisterResponse, error)
// 	ValidateToken(context.Context, *ValidateTokenRequest) (*BasicResponseCodeMessage, error)
// }

// Adapter implements the GRPCPort interface
type adapter struct {
	minipulsa.UnimplementedAuthenticationServiceServer
	api port.APIPort
}

// NewAdapter creates a new Adapter
func NewAdapter(api port.APIPort) adapter {
	return adapter{
		api: api,
	}
}

func (a adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	userServer := a
	grpcServer := grpc.NewServer()
	minipulsa.RegisterAuthenticationServiceServer(grpcServer, userServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}

func (a adapter) LoginRegister(context.Context, *minipulsa.LoginRegisterRequest) (*minipulsa.LoginRegisterResponse, error) {
	return nil, nil
}

func (a adapter) ValidateToken(context.Context, *minipulsa.ValidateTokenRequest) (*minipulsa.BasicResponseCodeMessage, error) {
	return nil, nil
}
