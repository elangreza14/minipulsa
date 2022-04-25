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
	GetWalletDetail(ctx context.Context, userID int64) (*entity.GetDBWallet, error)
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

func (ws walletService) GetWalletDetail(ctx context.Context, userID int64) (*entity.GetDBWallet, error) {
	grpcReq := &minipulsa.GetWalletDetailRequest{
		UserId: userID,
	}

	resGrpc, err := ws.WalletServiceClient.GetWalletDetail(ctx, grpcReq)
	if err != nil {
		ws.log.Logger.Error("LoginRegister ERR: ", err)
		return nil, err
	}

	history := []entity.DBHistoryWallet{}
	for i := 0; i < len(resGrpc.HistoryWallet); i++ {
		history = append(history, entity.DBHistoryWallet{
			LastAmount: resGrpc.HistoryWallet[i].LastAmount,
			OrderID:    resGrpc.HistoryWallet[i].OrderId,
			Date:       resGrpc.HistoryWallet[i].Date,
		})
	}
	res := &entity.GetDBWallet{
		Detail: entity.DBWallet{
			WalletID: resGrpc.Wallet.WalletId,
			UserID:   userID,
			OrderID:  resGrpc.Wallet.OrderId,
			Amount:   resGrpc.Wallet.Amount,
			Date:     resGrpc.Wallet.Date,
		},
		History: history,
	}
	return res, nil
}
