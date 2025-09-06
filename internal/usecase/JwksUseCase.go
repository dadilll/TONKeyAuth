package usecase

import (
	"TON/internal/dto"
	"crypto/rsa"
	"encoding/base64"
	"math/big"
)

type JWKSUseCase interface {
	GetJWKS() (*dto.JWKSResponseDTO, error)
	GetPublicKey() *rsa.PublicKey
}

type JWKSUseCaseImpl struct {
	KeyID string
	Pub   *rsa.PublicKey
}

func NewJWKSUseCase(keyID string, pub *rsa.PublicKey) JWKSUseCase {
	return &JWKSUseCaseImpl{
		KeyID: keyID,
		Pub:   pub,
	}
}

func (u *JWKSUseCaseImpl) GetJWKS() (*dto.JWKSResponseDTO, error) {
	n := base64.RawURLEncoding.EncodeToString(u.Pub.N.Bytes())
	e := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(u.Pub.E)).Bytes())

	jwk := dto.JWK{
		Kid: u.KeyID,
		Kty: "RSA",
		Alg: "RS256",
		Use: "sig",
		N:   n,
		E:   e,
	}

	return &dto.JWKSResponseDTO{
		Keys: []dto.JWK{jwk},
	}, nil
}

func (u *JWKSUseCaseImpl) GetPublicKey() *rsa.PublicKey {
	return u.Pub
}
