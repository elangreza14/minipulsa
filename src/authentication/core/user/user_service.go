package user

import (
	"context"
	"time"

	"github.com/elangreza14/minipulsa/authentication/entity"
	"github.com/elangreza14/minipulsa/authentication/token"
	"github.com/sirupsen/logrus"
)

// UserService is service layer that handle interaction between core and adapter
type UserService interface {
	LoginRegister(context.Context, entity.ReqPostPutUser) (string, *int64, error)
	ValidateToken(context.Context, string) (*int64, error)
}

type userService struct {
	userRepo UserRepository
	log      *logrus.Entry
	token    token.Maker
}

// NewUserService is generating new instance
func NewUserService(UserRepo UserRepository,
	log *logrus.Entry,
) UserService {
	tokenMaker, _ := token.NewJWTMaker("testing-base-testing-base-testing-base", time.Duration(time.Hour*6))
	return &userService{
		userRepo: UserRepo,
		log:      log,
		token:    tokenMaker,
	}
}

func (us *userService) LoginRegister(ctx context.Context, req entity.ReqPostPutUser) (string, *int64, error) {
	res, _ := us.userRepo.GetUserByEmail(ctx, req.Email)

	resToken, _ := us.token.CreateToken(req.Email)

	req.Token = resToken

	if res != nil {
		if res.Password != req.Password {
			err := entity.ErrPassWrong
			us.log.Error("LoginRegister ERROR: ", err)
			return "", nil, entity.ErrorGRPCPasswordWrong
		}

		err := us.userRepo.UpdateUser(ctx, req, res.UserID)
		if err != nil {
			us.log.Error("LoginRegister ERROR: ", err)
			return "", nil, entity.ErrorGRPCInternalServer
		}

		us.log.Info("LoginRegister SUCCESS: ")
		return req.Token, &res.UserID, nil
	}

	resCreate, err := us.userRepo.InsertUser(ctx, req)
	if err != nil {
		us.log.Error("LoginRegister ERROR: ", err)
		return "", nil, entity.ErrorGRPCInternalServer
	}

	us.log.Info("LoginRegister SUCCESS: ")
	return req.Token, resCreate, nil
}

func (us *userService) ValidateToken(ctx context.Context, token string) (*int64, error) {
	resToken, err := us.token.VerifyToken(token)
	if err != nil {
		us.log.Error("ValidateToken ERROR: ", err)
		return nil, entity.ErrorGRPCUnauthorized
	}

	res, err := us.userRepo.GetUserByEmail(ctx, resToken.Email)
	if err != nil {
		us.log.Error("ValidateToken ERROR: ", err)
		return nil, entity.ErrorGRPCInternalServer
	}

	if token != res.Token {
		err := entity.ErrorGRPCTokenWrong
		us.log.Error("ValidateToken ERROR: ", err)
		return nil, err
	}

	us.log.Info("ValidateToken SUCCESS: ")
	return &res.UserID, nil
}
