package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/dzniel/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	retrievedAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, newAccount.ID, retrievedAccount.ID)
	require.Equal(t, newAccount.Owner, retrievedAccount.Owner)
	require.Equal(t, newAccount.Balance, retrievedAccount.Balance)
	require.Equal(t, newAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: util.RandomBalance(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, newAccount.ID, updatedAccount.ID)
	require.Equal(t, newAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, newAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	retrievedAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, retrievedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
