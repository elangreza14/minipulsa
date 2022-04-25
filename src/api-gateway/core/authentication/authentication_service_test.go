package authentication

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/elangreza14/minipulsa/api-gateway/mock"
	"github.com/elangreza14/minipulsa/api-gateway/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_authenticationService_LoginRegister(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("err LoginRegister", func(t *testing.T) {
		authRepoMock := mock.NewMockAuthRepo(ctrl)
		as := authenticationService{
			AuthenticationServiceClient: authRepoMock,
			log:                         loggerserver.NewLogger(),
		}
		errBase := errors.New("err")

		base := entity.HTTPReqPostPutUser{
			Email:    util.RandomEmail(),
			Password: util.RandomString(10),
		}
		authRepoMock.EXPECT().LoginRegister(context.Background(), &minipulsa.LoginRegisterRequest{
			Email:    base.Email,
			Password: base.Password,
		}).Return(nil, errBase).AnyTimes()

		_, _, err := as.LoginRegister(context.Background(), base)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		authRepoMock := mock.NewMockAuthRepo(ctrl)
		as := authenticationService{
			AuthenticationServiceClient: authRepoMock,
			log:                         loggerserver.NewLogger(),
		}

		base := entity.HTTPReqPostPutUser{
			Email:    util.RandomEmail(),
			Password: util.RandomString(10),
		}

		res := &minipulsa.LoginRegisterResponse{}
		authRepoMock.EXPECT().LoginRegister(context.Background(), &minipulsa.LoginRegisterRequest{
			Email:    base.Email,
			Password: base.Password,
		}).Return(res, nil).AnyTimes()

		_, _, err := as.LoginRegister(context.Background(), base)
		require.NoError(t, err)
	})

}
