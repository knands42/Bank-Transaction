package factory

import (
	"github.com/caiofernandes00/Bank-Transaction.git/app/internal/token"
	"github.com/caiofernandes00/Bank-Transaction.git/app/internal/token/jwt"
	"github.com/caiofernandes00/Bank-Transaction.git/app/internal/token/paseto"
	"github.com/caiofernandes00/Bank-Transaction.git/app/internal/util"
)

func TokenFactory(config *util.Config) (token.Maker, error) {
	switch config.TokenType {
	case "jwt":
		return jwt.NewJWTMaker(config.TokenSymmetricKey)
	default:
		return paseto.NewPasetoMaker(config.TokenSymmetricKey)
	}
}
