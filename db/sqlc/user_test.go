package db

import (
	"context"
	"testing"
	"time"

	"github.com/dzniel/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := createRandomUser(t)
	retrievedUser, err := testQueries.GetUser(context.Background(), newUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	require.Equal(t, newUser.Username, retrievedUser.Username)
	require.Equal(t, newUser.HashedPassword, retrievedUser.HashedPassword)
	require.Equal(t, newUser.FullName, retrievedUser.FullName)
	require.Equal(t, newUser.Email, retrievedUser.Email)
	require.WithinDuration(t, newUser.PasswordChangedAt, retrievedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, newUser.CreatedAt, retrievedUser.CreatedAt, time.Second)
}
