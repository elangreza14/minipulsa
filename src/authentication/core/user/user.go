package user

import (
	"context"

	"github.com/elangreza14/minipulsa/authentication/entity"
)

type (
	// UserRepository handle all repository with this 6 method availabel
	// This Repository is not concern of any adapters. Can be postgres / redis/ mongo/ maria
	// This Repository also concern about handling all method availabel

	UserRepository interface {
		InsertUser(ctx context.Context, req entity.ReqPostPutUser) (*int64, error)
		UpdateUser(ctx context.Context, req entity.ReqPostPutUser, userID int64) error
		GetUser(ctx context.Context, userID int64) (*entity.DBUser, error)
		GetUserByEmail(ctx context.Context, email string) (*entity.DBUser, error)
	}

	// APIPort interface {
	// 	LoginRegister(ctx context.Context, in *minipulsa.LoginRegisterRequest, opts ...grpc.CallOption) (*minipulsa.LoginRegisterResponse, error)
	// 	ValidateToken(ctx context.Context, in *minipulsa.ValidateTokenRequest, opts ...grpc.CallOption) (*minipulsa.BasicResponseCodeMessage, error)
	// }
)
