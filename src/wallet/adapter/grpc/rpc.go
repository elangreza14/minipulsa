package grpc

import (
	"context"

	minipulsa "github.com/elangreza14/minipulsa/wallet/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/wallet/entity"
)

func (a adapter) GetWalletDetail(ctx context.Context, req *minipulsa.GetWalletDetailRequest) (*minipulsa.GetWalletDetailResponse, error) {
	res, err := a.api.GetWalletDetail(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	respHistory := []*minipulsa.HistoryWallet{}

	for i := 0; i < len(res.HistoryWallet); i++ {
		baseHistory := &minipulsa.HistoryWallet{
			LastAmount: res.HistoryWallet[i].LastAmount,
			OrderId:    res.HistoryWallet[i].OrderID.Int64,
			Date:       res.HistoryWallet[i].Date.Format(entity.YMDT),
		}
		respHistory = append(respHistory, baseHistory)
	}

	resp := &minipulsa.GetWalletDetailResponse{
		Wallet: &minipulsa.Wallet{
			WalletId: res.WalletID,
			UserId:   res.UserID,
			Amount:   res.Amount,
			OrderId:  res.OrderID,
			Date:     res.Date.Format(entity.YMDT),
		},
		HistoryWallet: respHistory,
	}

	return resp, nil
}

func (a adapter) UseWallet(ctx context.Context, req *minipulsa.UseWalletRequest) (*minipulsa.UseWalletResponse, error) {
	reqApi := &entity.ReqUseWallet{
		UserID:  req.UserId,
		Amount:  req.Amount,
		OrderID: req.OrderId,
	}

	wallet_history_id, err := a.api.UseWallet(ctx, *reqApi)
	if err != nil {
		return nil, err
	}

	resp := &minipulsa.UseWalletResponse{
		WalletHistoryId: *wallet_history_id,
	}

	return resp, nil

}
