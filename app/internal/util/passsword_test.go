package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	salt := RandomString(6)
	hashingConfig := NewHashingConfig(salt)

	hashedPassword1 := hashingConfig.HashPassword(password)
	require.NotEmpty(t, hashedPassword1)

	isEqual := hashingConfig.CheckPassword(password, hashedPassword1)
	require.True(t, isEqual)

	wrongPassword := RandomString(6)
	isEqual = hashingConfig.CheckPassword(wrongPassword, hashedPassword1)
	require.False(t, isEqual)

	hashedPassword2 := hashingConfig.HashPassword(password)
	require.NotEmpty(t, hashedPassword2)
	require.Equal(t, hashedPassword1, hashedPassword2)
}
