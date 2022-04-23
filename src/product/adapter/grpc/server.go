package grpc

import (
	"log"
	"net"

	minipulsa "github.com/elangreza14/minipulsa/product/adapter/grpc/minipulsa"
	port "github.com/elangreza14/minipulsa/product/port"

	grpc "google.golang.org/grpc"
)

type adapter struct {
	minipulsa.UnimplementedProductServiceServer
	api port.APIPort
}

// NewAdapter creates a new Adapter
func NewAdapter(api port.APIPort) adapter {
	return adapter{
		api: api,
	}
}

func (a adapter) Run() {
	log.Println("port 9001 ready to listen")
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen on port 9001: %v", err)
	}

	productServer := a
	grpcServer := grpc.NewServer()

	minipulsa.RegisterProductServiceServer(grpcServer, productServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9001: %v", err)
	}

}
