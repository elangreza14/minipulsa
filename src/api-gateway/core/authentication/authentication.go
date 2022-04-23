package authentication

import (
	"context"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"google.golang.org/grpc"
)

type (
	// UserRepository handle all repository with this 6 method availabel
	// This Repository is not concern of any adapters. Can be postgres / redis/ mongo/ maria
	// This Repository also concern about handling all method availabel

	UserRepository interface {
		LoginRegister(ctx context.Context, in *minipulsa.LoginRegisterRequest, opts ...grpc.CallOption) (*minipulsa.LoginRegisterResponse, error)
		ValidateToken(ctx context.Context, in *minipulsa.ValidateTokenRequest, opts ...grpc.CallOption) (*minipulsa.ValidateTokenResponse, error)
	}
)
