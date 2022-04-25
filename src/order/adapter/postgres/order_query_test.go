package postgres

import (
	"context"
	"testing"

	"github.com/elangreza14/minipulsa/order/entity"
	"github.com/elangreza14/minipulsa/order/util"
	"github.com/stretchr/testify/require"
)

func TestQueries_InsertOrder(t *testing.T) {
	res, err := testQueries.InsertOrder(context.Background(), entity.ReqCreateOrder{
		ProductID: util.RandomInt(1, 10000000),
		UserID:    util.RandomInt(1, 10000000),
	}, entity.OrderStatusPending)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestQueries_UpdateOrder(t *testing.T) {
	res, err := testQueries.InsertOrder(context.Background(), entity.ReqCreateOrder{
		ProductID: util.RandomInt(1, 10000000),
		UserID:    util.RandomInt(1, 10000000),
	}, entity.OrderStatusPending)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	resUpdate, err := testQueries.UpdateOrder(context.Background(), entity.ReqUpdateOrder{
		OrderID:         res.OrderID,
		Status:          string(entity.OrderStatusSUCCESS),
		WalletHistoryID: 11,
	})
	require.NoError(t, err)
	require.NotEmpty(t, resUpdate)

}

func TestQueries_GetOrder(t *testing.T) {
	res, err := testQueries.InsertOrder(context.Background(), entity.ReqCreateOrder{
		ProductID: util.RandomInt(1, 10000000),
		UserID:    util.RandomInt(1, 10000000),
	}, entity.OrderStatusPending)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	resGet, err := testQueries.GetOrder(context.Background(), res.OrderID)
	require.NoError(t, err)
	require.NotEmpty(t, resGet)
}
