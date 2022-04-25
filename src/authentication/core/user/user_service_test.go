package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/elangreza14/minipulsa/authentication/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/authentication/entity"
	"github.com/elangreza14/minipulsa/authentication/mock"
	"github.com/elangreza14/minipulsa/authentication/token"
	"github.com/elangreza14/minipulsa/authentication/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_userService_LoginRegister(t *testing.T) {
	crtl := gomock.NewController(t)
	tokenMaker, _ := token.NewJWTMaker("testing-base-testing-base-testing-base", time.Duration(time.Hour*6))

	t.Run("Error InsertUser", func(t *testing.T) {
		mockUserRepo := mock.NewMockUserRepository(crtl)

		us := &userService{
			userRepo: mockUserRepo,
			log:      loggerserver.NewLogger(),
			token:    tokenMaker,
		}
		email := util.RandomEmail()
		token, _ := tokenMaker.CreateToken(email)
		baseReq := entity.ReqPostPutUser{
			Email:    email,
			Password: util.RandomString(10),
			Token:    token,
		}

		errBase := errors.New("err")
		resGetEmail := &entity.DBUser{}
		mockUserRepo.EXPECT().GetUserByEmail(context.Background(), baseReq.Email).Return(resGetEmail, nil).AnyTimes()
		mockUserRepo.EXPECT().InsertUser(context.Background(), baseReq).Return(nil, errBase).AnyTimes()

		_, _, err := us.LoginRegister(context.Background(), baseReq)
		require.Error(t, err)

	})

	t.Run("Error Password Not same", func(t *testing.T) {
		mockUserRepo := mock.NewMockUserRepository(crtl)

		us := &userService{
			userRepo: mockUserRepo,
			log:      loggerserver.NewLogger(),
			token:    tokenMaker,
		}
		email := util.RandomEmail()
		token, _ := tokenMaker.CreateToken(email)
		baseReq := entity.ReqPostPutUser{
			Email:    email,
			Password: util.RandomString(10),
			Token:    token,
		}

		resGetEmail := &entity.DBUser{}
		mockUserRepo.EXPECT().GetUserByEmail(context.Background(), baseReq.Email).Return(resGetEmail, nil).AnyTimes()

		_, _, err := us.LoginRegister(context.Background(), baseReq)
		require.Error(t, err)

	})

	t.Run("Success", func(t *testing.T) {
		mockUserRepo := mock.NewMockUserRepository(crtl)

		us := &userService{
			userRepo: mockUserRepo,
			log:      loggerserver.NewLogger(),
			token:    tokenMaker,
		}
		email := util.RandomEmail()
		token, _ := tokenMaker.CreateToken(email)
		baseReq := entity.ReqPostPutUser{
			Email:    email,
			Password: util.RandomString(10),
			Token:    token,
		}

		resGetEmail := &entity.DBUser{}
		mockUserRepo.EXPECT().GetUserByEmail(context.Background(), baseReq.Email).Return(resGetEmail, nil).AnyTimes()
		mockUserRepo.EXPECT().UpdateUser(context.Background(), baseReq, resGetEmail.UserID).Return(nil).AnyTimes()

		_, _, err := us.LoginRegister(context.Background(), baseReq)
		require.Error(t, err)

	})
}

func Test_userService_ValidateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	tokenMaker, _ := token.NewJWTMaker("testing-base-testing-base-testing-base", time.Duration(time.Hour*6))

	t.Run("error token", func(t *testing.T) {
		userRepoMock := mock.NewMockUserRepository(ctrl)
		us := &userService{
			userRepo: userRepoMock,
			log:      loggerserver.NewLogger(),
			token:    tokenMaker,
		}
		res := &entity.DBUser{}
		userRepoMock.EXPECT().GetUserByEmail(context.Background(), gomock.Any().String()).Return(res, nil).AnyTimes()
		_, err := us.ValidateToken(context.Background(), gomock.Any().String())

		require.Error(t, err)

	})

	t.Run("err GetUserByEmail", func(t *testing.T) {
		userRepoMock := mock.NewMockUserRepository(ctrl)
		us := &userService{
			userRepo: userRepoMock,
			log:      loggerserver.NewLogger(),
			token:    tokenMaker,
		}
		errBase := errors.New("err")
		baseString := gomock.Any().String()
		userRepoMock.EXPECT().GetUserByEmail(context.Background(), baseString).Return(nil, errBase).AnyTimes()

		_, err := us.ValidateToken(context.Background(), baseString)

		require.Error(t, err)

	})

}
