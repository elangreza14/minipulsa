package grpc

import (
	"context"

	minipulsa "github.com/elangreza14/minipulsa/wallet/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/wallet/entity"
)

func (a adapter) GetWalletDetail(ctx context.Context, req *minipulsa.GetWalletDetailRequest) (*minipulsa.GetWalletDetailResponse, error) {
	return nil, nil
}

func (a adapter) UseWallet(ctx context.Context, req *minipulsa.UseWalletRequest) (*minipulsa.BasicResponseCodeMessage, error) {
	reqApi := &entity.ReqUseWallet{
		UserID: req.UserId,
		Amount: req.Amount,
	}

	err := a.api.UseWallet(ctx, *reqApi)
	if err != nil {
		return nil, err
	}

	resp := &minipulsa.BasicResponseCodeMessage{
		Message: "OK",
	}
	return resp, nil

}
