package user

import (
	"github.com/sirupsen/logrus"
)

// UserService is service layer that handle interaction between core and adapter
type UserService interface {
	// LoginRegister(context.Context, *grpc.LoginRegisterRequest) (*grpc.LoginRegisterResponse, error)
	// ValidateToken(context.Context, *grpc.ValidateTokenRequest) (*grpc.BasicResponseCodeMessage, error)
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

// func (us *userService) LoginRegister(context.Context, *grpc.LoginRegisterRequest) (*grpc.LoginRegisterResponse, error) {
// 	return nil, nil
// }

// func (us *userService) ValidateToken(context.Context, *grpc.ValidateTokenRequest) (*grpc.BasicResponseCodeMessage, error) {
// 	return nil, nil
// }
