package util

import (
	"golang.org/x/crypto/argon2"
)

type HashingConfig struct {
	Salt      string
	Time      uint32
	Memory    uint32
	CPUNumber uint8
	KeyLength uint32
}

type HashOptions = func(*HashingConfig)

func defaultConfig() *HashingConfig {
	return &HashingConfig{
		Time:      1,
		Memory:    64 * 1024,
		CPUNumber: 4,
		KeyLength: 32,
	}
}

func NewHashingConfig(salt string, options ...HashOptions) *HashingConfig {
	config := defaultConfig()
	config.Salt = salt

	for _, option := range options {
		option(config)
	}

	return config
}

func (h *HashingConfig) HashPassword(password string) []byte {
	hashedPassword := argon2.IDKey([]byte(password), []byte(h.Salt), h.Time, h.Memory, h.CPUNumber, h.KeyLength)

	return hashedPassword
}

func (h *HashingConfig) CheckPassword(password string, hashedPassword []byte) bool {
	return string(h.HashPassword(password)) == string(hashedPassword)
}
