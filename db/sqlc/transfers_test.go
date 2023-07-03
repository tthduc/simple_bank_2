package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTransfers(t *testing.T) Transfer {
	account1, err := testQueries.GetAccount(context.Background(), 1)
	account2, err := testQueries.GetAccount(context.Background(), 2)
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.NotEmpty(t, account2)

	arg := CreateTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        100,
	}

	transfer, err := testQueries.CreateTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfers(t)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfers(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, account := range transfers {
		require.NotEmpty(t, account)
	}
}
