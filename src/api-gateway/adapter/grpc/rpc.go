package grpc

import (
	"context"

	minipulsa "github.com/elangreza14/minipulsa/authentication/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/authentication/entity"
)

func (a adapter) LoginRegister(ctx context.Context, req *minipulsa.LoginRegisterRequest) (*minipulsa.LoginRegisterResponse, error) {
	reqBase := entity.ReqPostPutUser{
		Email:    req.Email,
		Password: req.Password,
	}

	token, id, err := a.api.LoginRegister(ctx, reqBase)
	if err != nil {
		return nil, err
	}

	res := &minipulsa.LoginRegisterResponse{
		Token:  token,
		UserId: *id,
	}

	return res, nil
}

func (a adapter) ValidateToken(ctx context.Context, req *minipulsa.ValidateTokenRequest) (*minipulsa.ValidateTokenResponse, error) {
	userID, err := a.api.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	res := &minipulsa.ValidateTokenResponse{
		UserId: *userID,
	}

	return res, nil
}
