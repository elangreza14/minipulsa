package postgres

import (
	"context"
	"testing"

	"github.com/elangreza14/minipulsa/authentication/entity"
	"github.com/elangreza14/minipulsa/authentication/util"
	"github.com/stretchr/testify/require"
)

func TestQueries_InsertUser(t *testing.T) {
	arg := entity.ReqPostPutUser{
		Email:    util.RandomEmail(),
		Password: util.RandomString(10),
		Token:    util.RandomString(10),
	}
	res, err := testQueries.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestQueries_UpdateUser(t *testing.T) {
	arg := entity.ReqPostPutUser{
		Email:    util.RandomEmail(),
		Password: util.RandomString(10),
		Token:    util.RandomString(10),
	}
	res, err := testQueries.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, res)
	err = testQueries.UpdateUser(context.Background(), arg, *res)
	require.NoError(t, err)
}

func TestQueries_GetUser(t *testing.T) {

	arg := entity.ReqPostPutUser{
		Email:    util.RandomEmail(),
		Password: util.RandomString(10),
		Token:    util.RandomString(10),
	}
	res, err := testQueries.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, res)

	resp, err := testQueries.GetUser(context.Background(), *res)
	compare := &entity.DBUser{
		UserID:   *res,
		Email:    resp.Email,
		Password: resp.Password,
		Token:    resp.Token,
	}

	require.NoError(t, err)
	require.NotEmpty(t, resp)
	require.Equal(t, resp, compare)
}
