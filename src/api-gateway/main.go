package main

import (
	"log"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/httpserver"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/api-gateway/core"
	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"github.com/elangreza14/minipulsa/api-gateway/core/product"
	"google.golang.org/grpc"
)

func main() {
	logBase := loggerserver.NewLogger()

	connAuth, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer connAuth.Close()
	authGrpcService := minipulsa.NewAuthenticationServiceClient(connAuth)

	connProduct, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer connProduct.Close()
	productGrpcService := minipulsa.NewProductServiceClient(connProduct)

	UserService := authentication.NewAuthenticationService(logBase, authGrpcService)
	ProductService := product.NewProductService(logBase, productGrpcService)

	core := core.NewBaseApp(logBase, UserService, ProductService)

	httpserver.HttpServer(core)
}
