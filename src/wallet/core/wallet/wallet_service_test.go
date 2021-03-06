package wallet

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/minipulsa/wallet/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/wallet/entity"
	"github.com/elangreza14/minipulsa/wallet/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_walletService_UseWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("Error GetWalletByUserID", func(t *testing.T) {

		mockWalletRepo := mock.NewMockWalletRepository(ctrl)
		logger := loggerserver.NewLogger()

		ws := walletService{
			walletRepo: mockWalletRepo,
			log:        logger,
		}
		baseReq := entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		}

		errBase := errors.New("err")

		mockWalletRepo.EXPECT().GetWalletByUserID(context.Background(), baseReq.UserID).Return(nil, errBase).Times(1).AnyTimes()
		mockWalletRepo.EXPECT().InsertWallet(context.Background(), baseReq).Return(errBase).AnyTimes()

		_, err := ws.UseWallet(context.Background(), entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		})
		require.Error(t, err)
	})
	t.Run("Error UpdateWalletByUserID 1", func(t *testing.T) {
		mockWalletRepo := mock.NewMockWalletRepository(ctrl)
		logger := loggerserver.NewLogger()

		ws := walletService{
			walletRepo: mockWalletRepo,
			log:        logger,
		}
		base := &entity.DBWallet{}
		baseReq := entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		}
		errBase := errors.New("err")

		mockWalletRepo.EXPECT().GetWalletByUserID(context.Background(), baseReq.UserID).Return(base, nil).AnyTimes()
		mockWalletRepo.EXPECT().UpdateWalletByUserID(context.Background(), baseReq).Return(errBase).AnyTimes()
		_, err := ws.UseWallet(context.Background(), entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		})
		require.Error(t, err)
	})

	t.Run("Error GetWalletHistoryByReqUseWallet 2", func(t *testing.T) {

		mockWalletRepo := mock.NewMockWalletRepository(ctrl)
		logger := loggerserver.NewLogger()

		ws := walletService{
			walletRepo: mockWalletRepo,
			log:        logger,
		}
		baseReq := entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		}

		errBase := errors.New("err")

		mockWalletRepo.EXPECT().GetWalletByUserID(context.Background(), baseReq.UserID).Return(nil, errBase).Times(1).AnyTimes()
		mockWalletRepo.EXPECT().InsertWallet(context.Background(), baseReq).Return(nil).AnyTimes()
		mockWalletRepo.EXPECT().UpdateWalletByUserID(context.Background(), baseReq).Return(nil).AnyTimes()
		mockWalletRepo.EXPECT().GetWalletHistoryByReqUseWallet(context.Background(), baseReq).Return(nil, errBase).AnyTimes()

		_, err := ws.UseWallet(context.Background(), entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		})
		require.Error(t, err)
	})

	t.Run("Success", func(t *testing.T) {

		mockWalletRepo := mock.NewMockWalletRepository(ctrl)
		logger := loggerserver.NewLogger()

		ws := walletService{
			walletRepo: mockWalletRepo,
			log:        logger,
		}
		base := &entity.DBWallet{}
		baseReq := entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		}

		baseHistoryRes := &entity.DBWalletHistoriesDetail{}
		mockWalletRepo.EXPECT().GetWalletByUserID(context.Background(), baseReq.UserID).Return(base, nil).AnyTimes()
		mockWalletRepo.EXPECT().UpdateWalletByUserID(context.Background(), baseReq).Return(nil).AnyTimes()
		mockWalletRepo.EXPECT().GetWalletHistoryByReqUseWallet(context.Background(), baseReq).Return(baseHistoryRes, nil).AnyTimes()
		_, err := ws.UseWallet(context.Background(), entity.ReqUseWallet{
			UserID:  1,
			Amount:  1,
			OrderID: 1,
		})
		require.NoError(t, err)
	})
}

func Test_walletService_GetWalletDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Error GetWalletByUserID", func(t *testing.T) {
		mockWalletRepo := mock.NewMockWalletRepository(ctrl)
		logger := loggerserver.NewLogger()

		ws := walletService{
			walletRepo: mockWalletRepo,
			log:        logger,
		}
		userID := 1
		errBase := errors.New("err")
		mockWalletRepo.EXPECT().GetWalletByUserID(context.Background(), int64(userID)).Return(nil, errBase).AnyTimes()
		_, err := ws.GetWalletDetail(context.Background(), int64(userID))
		require.Error(t, err)
	})

	t.Run("Error GetWalletHistories", func(t *testing.T) {
		mockWalletRepo := mock.NewMockWalletRepository(ctrl)
		logger := loggerserver.NewLogger()

		ws := walletService{
			walletRepo: mockWalletRepo,
			log:        logger,
		}

		userID := 1
		base := &entity.DBWallet{}
		errBase := errors.New("err")
		mockWalletRepo.EXPECT().GetWalletByUserID(context.Background(), int64(userID)).Return(base, nil).AnyTimes()
		mockWalletRepo.EXPECT().GetWalletHistories(context.Background(), int64(userID)).Return(nil, errBase).AnyTimes()
		_, err := ws.GetWalletDetail(context.Background(), int64(userID))
		require.Error(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		mockWalletRepo := mock.NewMockWalletRepository(ctrl)
		logger := loggerserver.NewLogger()

		ws := walletService{
			walletRepo: mockWalletRepo,
			log:        logger,
		}

		userID := 1
		base := &entity.DBWallet{}
		baseResHistory := &[]entity.DBWalletHistories{}
		mockWalletRepo.EXPECT().GetWalletByUserID(context.Background(), int64(userID)).Return(base, nil).AnyTimes()
		mockWalletRepo.EXPECT().GetWalletHistories(context.Background(), int64(userID)).Return(baseResHistory, nil).AnyTimes()
		_, err := ws.GetWalletDetail(context.Background(), int64(userID))
		require.NoError(t, err)
	})
}
