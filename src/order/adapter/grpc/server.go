package grpc

import (
	"log"
	"net"

	minipulsa "github.com/elangreza14/minipulsa/order/adapter/grpc/minipulsa"
	port "github.com/elangreza14/minipulsa/order/port"

	grpc "google.golang.org/grpc"
)

type adapter struct {
	minipulsa.UnimplementedOrderServiceServer
	api port.ApiPort
}

// NewAdapter creates a new Adapter
func NewAdapter(api port.ApiPort) adapter {
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

	orderServer := a
	grpcServer := grpc.NewServer()

	minipulsa.RegisterOrderServiceServer(grpcServer, orderServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9001: %v", err)
	}

}
