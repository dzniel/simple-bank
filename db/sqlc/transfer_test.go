package db

import (
	"context"
	"testing"
	"time"

	"github.com/dzniel/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomBalance(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
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
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	CreateRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	newTransfer := CreateRandomTransfer(t, account1, account2)

	retrievedTransfer, err := testQueries.GetTransfer(context.Background(), newTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedTransfer)

	require.Equal(t, newTransfer.ID, retrievedTransfer.ID)
	require.Equal(t, newTransfer.FromAccountID, retrievedTransfer.FromAccountID)
	require.Equal(t, newTransfer.ToAccountID, retrievedTransfer.ToAccountID)
	require.Equal(t, newTransfer.Amount, retrievedTransfer.Amount)
	require.WithinDuration(t, newTransfer.CreatedAt, retrievedTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, account1, account2)
		CreateRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, arg.FromAccountID == transfer.FromAccountID || arg.ToAccountID == transfer.ToAccountID)
	}
}
