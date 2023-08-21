package gapi

import (
	"testing"
	"time"

	db "github.com/dzniel/simple-bank/db/sqlc"
	"github.com/dzniel/simple-bank/util"
	"github.com/dzniel/simple-bank/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
