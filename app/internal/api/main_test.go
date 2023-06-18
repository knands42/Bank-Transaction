package api

import (
	"os"
	"testing"
	"time"

	db "github.com/caiofernandes00/Bank-Transaction.git/app/internal/db/sqlc"
	"github.com/caiofernandes00/Bank-Transaction.git/app/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var SALT = util.RandomString(6)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	hashingConfig := util.NewHashingConfig(SALT)

	server, err := NewServer(config, store, hashingConfig)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
