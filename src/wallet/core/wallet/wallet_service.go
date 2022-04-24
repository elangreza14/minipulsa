package wallet

import (
	"context"
	"database/sql"

	"github.com/elangreza14/minipulsa/wallet/entity"
	"github.com/sirupsen/logrus"
)

// WalletService is service layer that handle interaction between core and adapter
type WalletService interface {
	UseWallet(ctx context.Context, req entity.ReqUseWallet) (*int64, error)
	GetWalletDetail(ctx context.Context, userId int64) (*entity.GetDetailWallet, error)
}

type walletService struct {
	walletRepo WalletRepository
	log        *logrus.Entry
}

// NewWalletService is generating new instance
func NewWalletService(WalletRepo WalletRepository,
	log *logrus.Entry,
) WalletService {
	return &walletService{
		walletRepo: WalletRepo,
		log:        log,
	}
}

func (ws *walletService) UseWallet(ctx context.Context, req entity.ReqUseWallet) (*int64, error) {
	wallet, _ := ws.walletRepo.GetWalletByUserID(ctx, req.UserID)

	if wallet == nil {
		err := ws.walletRepo.InsertWallet(ctx, req)
		if err != nil {
			ws.log.Error("UseWallet ERROR: ", err)
			return nil, entity.ErrorGRPCInternalServer
		}
	} else {
		err := ws.walletRepo.UpdateWalletByUserID(ctx, req)
		if err != nil {
			ws.log.Error("UseWallet ERROR: ", err)
			return nil, entity.ErrorGRPCInternalServer
		}
	}

	walletHistory, err := ws.walletRepo.GetWalletHistoryByReqUseWallet(ctx, req)
	if err != nil {
		ws.log.Error("UseWallet ERROR: ", err)
		return nil, entity.ErrorGRPCInternalServer
	}

	ws.log.Info("UseWallet SUCCESS: ")
	return &walletHistory.WalletHistoryID, nil
}

func (ws *walletService) GetWalletDetail(ctx context.Context, userId int64) (*entity.GetDetailWallet, error) {
	wallet, err := ws.walletRepo.GetWalletByUserID(ctx, userId)

	if err != nil {
		ws.log.Error("GetWalletDetail ERROR: ", err)
		if err == sql.ErrNoRows {
			return nil, entity.ErrorGRPCNotFound
		}
		return nil, entity.ErrorGRPCInternalServer
	}

	walletHistory, err := ws.walletRepo.GetWalletHistories(ctx, userId)

	if err != nil {
		ws.log.Error("GetWalletDetail ERROR: ", err)
		if err == sql.ErrNoRows {
			return nil, entity.ErrorGRPCNotFound
		}
		return nil, entity.ErrorGRPCInternalServer
	}

	res := &entity.GetDetailWallet{
		WalletID:      wallet.WalletID,
		UserID:        userId,
		Amount:        wallet.Amount,
		OrderID:       wallet.OrderID.Int64,
		Date:          wallet.Date,
		HistoryWallet: *walletHistory,
	}

	ws.log.Info("GetWalletDetail SUCCESS: ")
	return res, nil
}
