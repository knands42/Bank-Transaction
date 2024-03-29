package paseto

import (
	"fmt"
	"time"

	token_maker "github.com/caiofernandes00/Bank-Transaction.git/app/internal/token"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (token_maker.Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (*token_maker.Payload, string, error) {
	payload, err := token_maker.NewPayload(username, duration)
	if err != nil {
		return nil, "", err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return payload, token, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*token_maker.Payload, error) {
	payload := &token_maker.Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, token_maker.ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
