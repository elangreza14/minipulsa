package postgres

import (
	"context"
	"testing"

	"github.com/elangreza14/minipulsa/wallet/entity"
	"github.com/elangreza14/minipulsa/wallet/util"
	"github.com/stretchr/testify/require"
)

func TestQueries_InsertWallet(t *testing.T) {
	req := entity.ReqUseWallet{
		UserID:  util.RandomInt(10, 100000),
		Amount:  10000000,
		OrderID: util.RandomInt(10, 100000),
	}
	err := testQueries.InsertWallet(context.Background(), req)
	require.NoError(t, err)
}

func TestQueries_GetWalletByUserID(t *testing.T) {
	req := entity.ReqUseWallet{
		UserID:  util.RandomInt(10, 100000),
		Amount:  10000000,
		OrderID: util.RandomInt(10, 100000),
	}
	err := testQueries.InsertWallet(context.Background(), req)
	require.NoError(t, err)

	res, err := testQueries.GetWalletByUserID(context.Background(), req.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, res.UserID, res.UserID)
}

func TestQueries_UpdateWalletByUserID(t *testing.T) {
	req := entity.ReqUseWallet{
		UserID:  util.RandomInt(10, 100000),
		Amount:  10000000,
		OrderID: util.RandomInt(10, 100000),
	}
	err := testQueries.InsertWallet(context.Background(), req)
	require.NoError(t, err)

	res, err := testQueries.GetWalletByUserID(context.Background(), req.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, res.UserID, res.UserID)

	reqUpdate := entity.ReqUseWallet{
		UserID:  res.UserID,
		Amount:  20000000,
		OrderID: res.OrderID.Int64,
	}

	err = testQueries.UpdateWalletByUserID(context.Background(), reqUpdate)
	require.NoError(t, err)

}

func TestQueries_GetWalletHistories(t *testing.T) {
	req := entity.ReqUseWallet{
		UserID:  util.RandomInt(10, 100000),
		Amount:  10000000,
		OrderID: util.RandomInt(10, 100000),
	}
	err := testQueries.InsertWallet(context.Background(), req)
	require.NoError(t, err)

	res, err := testQueries.GetWalletHistories(context.Background(), req.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

}

func TestQueries_GetWalletHistoryByReqUseWallet(t *testing.T) {
	req := entity.ReqUseWallet{
		UserID:  util.RandomInt(10, 100000),
		Amount:  10000000,
		OrderID: util.RandomInt(10, 100000),
	}
	err := testQueries.InsertWallet(context.Background(), req)
	require.NoError(t, err)

	res, err := testQueries.GetWalletHistoryByReqUseWallet(context.Background(), req)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
