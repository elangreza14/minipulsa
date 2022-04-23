package user

import (
	"context"

	"github.com/elangreza14/minipulsa/authentication/entity"
	"github.com/sirupsen/logrus"
)

// UserService is service layer that handle interaction between core and adapter
type UserService interface {
	LoginRegisterUser(ctx context.Context, arg entity.ReqPostPutUser) (*int64, error)
	ValidateUser(ctx context.Context, token string) error
}

type userService struct {
	userRepo UserRepository
	log      *logrus.Entry
}

// NewUserService is generating new instance
func NewUserService(UserRepo UserRepository,
	log *logrus.Entry,
) UserService {
	return &userService{
		userRepo: UserRepo,
		log:      log,
	}
}

// AddUser is adding user method
func (ws *userService) LoginRegisterUser(ctx context.Context, arg entity.ReqPostPutUser) (*int64, error) {

	return nil, nil
}

// ValidateUser is update user that already exist
func (ws *userService) ValidateUser(ctx context.Context, token string) error {

	return nil
}
