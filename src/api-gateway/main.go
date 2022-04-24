package main

import (
	"log"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/httpserver"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/api-gateway/core"
	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"github.com/elangreza14/minipulsa/api-gateway/core/order"
	"github.com/elangreza14/minipulsa/api-gateway/core/product"
	"github.com/elangreza14/minipulsa/api-gateway/core/wallet"
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
	UserService := authentication.NewAuthenticationService(logBase, authGrpcService)

	connProduct, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer connProduct.Close()
	productGrpcService := minipulsa.NewProductServiceClient(connProduct)
	ProductService := product.NewProductService(logBase, productGrpcService)

	connWallet, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer connWallet.Close()
	walletGrpcService := minipulsa.NewWalletServiceClient(connWallet)
	WalletService := wallet.NewWalletService(logBase, walletGrpcService)

	connOrder, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer connOrder.Close()
	orderGrpcService := minipulsa.NewOrderServiceClient(connOrder)
	OrderService := order.NewOrderService(logBase, orderGrpcService, walletGrpcService, productGrpcService)

	core := core.NewBaseApp(logBase, UserService, ProductService, WalletService, OrderService)

	httpserver.HttpServer(core)
}
