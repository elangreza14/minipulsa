package authentication

import (
	"context"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/sirupsen/logrus"
)

// AuthenticationService is service layer that handle interaction between core and adapter
type AuthenticationService interface {
	LoginRegister(context.Context, entity.HTTPReqPostPutUser) (string, *int64, error)
}

type authenticationService struct {
	AuthenticationServiceClient minipulsa.AuthenticationServiceClient
	log                         *logrus.Entry
}

// NewAuthenticationService is generating new instance
func NewAuthenticationService(
	log *logrus.Entry,
	AuthenticationServiceClient minipulsa.AuthenticationServiceClient,
) AuthenticationService {
	return &authenticationService{
		AuthenticationServiceClient: AuthenticationServiceClient,
		log:                         log,
	}
}

func (us *authenticationService) LoginRegister(ctx context.Context, req entity.HTTPReqPostPutUser) (string, *int64, error) {
	grpcReq := &minipulsa.LoginRegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := us.AuthenticationServiceClient.LoginRegister(ctx, grpcReq)
	if err != nil {
		us.log.Logger.Error("LoginRegister ERR: ", err)
		return "", nil, err
	}
	// us.log.Logger.Info("res ", res)
	// us.log.Logger.Info("err ", err)

	return res.Token, &res.UserId, nil
}
