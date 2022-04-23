package product

import (
	"context"

	"github.com/elangreza14/minipulsa/product/entity"
)

type (
	// ProductRepository handle all repository with this 6 method availabel
	// This Repository is not concern of any adapters. Can be postgres / redis/ mongo/ maria
	// This Repository also concern about handling all method availabel

	ProductRepository interface {
		GetProductByID(ctx context.Context, productID int64) (*entity.DBProduct, error)
		GetProducts(ctx context.Context) (*[]entity.DBProduct, error)
	}

	// APIPort interface {
	// 	LoginRegister(ctx context.Context, in *minipulsa.LoginRegisterRequest, opts ...grpc.CallOption) (*minipulsa.LoginRegisterResponse, error)
	// 	ValidateToken(ctx context.Context, in *minipulsa.ValidateTokenRequest, opts ...grpc.CallOption) (*minipulsa.BasicResponseCodeMessage, error)
	// }
)
