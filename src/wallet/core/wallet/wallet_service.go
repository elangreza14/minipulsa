package wallet

import (
	"context"
	"fmt"

	"github.com/elangreza14/minipulsa/wallet/entity"
	"github.com/sirupsen/logrus"
)

// WalletService is service layer that handle interaction between core and adapter
type WalletService interface {
	UseWallet(ctx context.Context, req entity.ReqUseWallet) error
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

func (ws *walletService) UseWallet(ctx context.Context, req entity.ReqUseWallet) error {
	wallet, _ := ws.walletRepo.GetWalletByUserID(ctx, req.UserID)

	if wallet == nil {
		err := ws.walletRepo.InsertWallet(ctx, req.Amount, req.UserID)
		if err != nil {
			fmt.Println(err)
			ws.log.Error("UseWallet ERROR: ", err)
			return entity.ErrorGRPCInternalServer
		}

		return nil
	}

	err := ws.walletRepo.UpdateWalletByUserID(ctx, req.Amount, req.UserID)
	if err != nil {
		ws.log.Error("UseWallet ERROR: ", err)
		return entity.ErrorGRPCInternalServer
	}
	return nil
}
