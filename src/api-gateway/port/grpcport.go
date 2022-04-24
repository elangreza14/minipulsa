package port

import (
	"context"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"google.golang.org/grpc"
)

type (
	AuthRepo interface {
		LoginRegister(ctx context.Context, in *minipulsa.LoginRegisterRequest, opts ...grpc.CallOption) (*minipulsa.LoginRegisterResponse, error)
		ValidateToken(ctx context.Context, in *minipulsa.ValidateTokenRequest, opts ...grpc.CallOption) (*minipulsa.ValidateTokenResponse, error)
	}

	ProductRepo interface {
		GetProducts(ctx context.Context, in *minipulsa.Empty, opts ...grpc.CallOption) (*minipulsa.GetProductsResponse, error)
		GetProduct(ctx context.Context, in *minipulsa.GetProductRequest, opts ...grpc.CallOption) (*minipulsa.GetProductResponse, error)
	}

	WalletRepo interface {
		GetWalletDetail(ctx context.Context, in *minipulsa.GetWalletDetailRequest, opts ...grpc.CallOption) (*minipulsa.GetWalletDetailResponse, error)
		UseWallet(ctx context.Context, in *minipulsa.UseWalletRequest, opts ...grpc.CallOption) (*minipulsa.UseWalletResponse, error)
	}

	OrderRepo interface {
		CreateOrder(ctx context.Context, in *minipulsa.CreateOrderRequest, opts ...grpc.CallOption) (*minipulsa.CreateORUpdateOrderResponse, error)
		UpdateOrder(ctx context.Context, in *minipulsa.UpdateOrderRequest, opts ...grpc.CallOption) (*minipulsa.CreateORUpdateOrderResponse, error)
	}
)
