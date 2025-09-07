package usecase

import (
	"TON/internal/dto"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	CreateToken(req dto.TokenRequestDTO) (*dto.TokenResponseDTO, error)
}

type TokenUseCaseImpl struct {
	Issuer  string
	TTL     time.Duration
	PrivKey *rsa.PrivateKey
}

func NewTokenUseCase(issuer string, ttl time.Duration, priv *rsa.PrivateKey) TokenUseCase {
	return &TokenUseCaseImpl{
		Issuer:  issuer,
		TTL:     ttl,
		PrivKey: priv,
	}
}

func (u *TokenUseCaseImpl) CreateToken(req dto.TokenRequestDTO) (*dto.TokenResponseDTO, error) {
	tokenID, err := generateRandomString(16)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"jti": tokenID,
		"iss": u.Issuer,
		"exp": time.Now().Add(u.TTL).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenStr, err := token.SignedString(u.PrivKey)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponseDTO{
		JWT: tokenStr,
	}, nil
}
