package wallet

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/elangreza14/minipulsa/api-gateway/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_walletService_UseWallet(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("err UseWallet", func(t *testing.T) {
		walletRepoMock := mock.NewMockWalletRepo(ctrl)
		ps := walletService{
			WalletServiceClient: walletRepoMock,
			log:                 loggerserver.NewLogger(),
		}
		errBase := errors.New("err")

		userID := 1
		grpcReq := &minipulsa.UseWalletRequest{
			UserId:  int64(userID),
			Amount:  0,
			OrderId: 0,
		}
		req := entity.HTTPReqPostUseWallet{}
		walletRepoMock.EXPECT().UseWallet(context.Background(), grpcReq).Return(nil, errBase).AnyTimes()

		err := ps.UseWallet(context.Background(), req, int64(userID))
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		walletRepoMock := mock.NewMockWalletRepo(ctrl)
		ps := walletService{
			WalletServiceClient: walletRepoMock,
			log:                 loggerserver.NewLogger(),
		}
		userID := 1
		res := &minipulsa.UseWalletResponse{}
		grpcReq := &minipulsa.UseWalletRequest{
			UserId:  int64(userID),
			Amount:  0,
			OrderId: 0,
		}
		req := entity.HTTPReqPostUseWallet{}
		walletRepoMock.EXPECT().UseWallet(context.Background(), grpcReq).Return(res, nil).AnyTimes()

		err := ps.UseWallet(context.Background(), req, int64(userID))
		require.NoError(t, err)
	})
}
