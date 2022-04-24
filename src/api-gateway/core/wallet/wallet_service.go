package wallet

import (
	"context"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/elangreza14/minipulsa/api-gateway/port"
	"github.com/sirupsen/logrus"
)

// WalletService is service layer that handle interaction between core and adapter
type WalletService interface {
	UseWallet(ctx context.Context, req entity.HTTPReqPostUseWallet, userID int64) error
}

type walletService struct {
	WalletServiceClient port.WalletRepo
	log                 *logrus.Entry
}

// NewWalletService is generating new instance
func NewWalletService(
	log *logrus.Entry,
	WalletServiceClient port.WalletRepo,

) WalletService {
	return &walletService{
		WalletServiceClient: WalletServiceClient,
		log:                 log,
	}
}

func (ws walletService) UseWallet(ctx context.Context, req entity.HTTPReqPostUseWallet, userID int64) error {
	grpcReq := &minipulsa.UseWalletRequest{
		UserId: userID,
		Amount: req.Amount,
	}

	_, err := ws.WalletServiceClient.UseWallet(ctx, grpcReq)
	if err != nil {
		ws.log.Logger.Error("LoginRegister ERR: ", err)
		return err
	}

	return nil
}
