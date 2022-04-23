package main

import (
	"log"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/httpserver"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"google.golang.org/grpc"
)

func main() {
	logBase := loggerserver.NewLogger()

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	authgrpcService := minipulsa.NewAuthenticationServiceClient(conn)

	UserService := authentication.NewAuthenticationService(logBase, authgrpcService)
	// appBase := core.BaseApp{
	// 	Ws:  UserService,
	// 	Log: logBase,
	// }

	httpserver.HttpServer(UserService, logBase)
}
