package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueries_GetProductByID(t *testing.T) {
	res, err := testQueries.GetProductByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
func TestQueries_GetProducts(t *testing.T) {
	res, err := testQueries.GetProducts(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
