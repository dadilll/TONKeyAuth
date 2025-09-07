package usecase

import (
	"TON/internal/dto"
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type TokenVerifyUseCase interface {
	VerifyToken(req dto.VerifyTokenRequestDTO, pub *rsa.PublicKey) (*dto.VerifyTokenResponseDTO, error)
}

type TokenVerifyUseCaseImpl struct{}

func NewTokenVerifyUseCase() TokenVerifyUseCase {
	return &TokenVerifyUseCaseImpl{}
}

func (u *TokenVerifyUseCaseImpl) VerifyToken(req dto.VerifyTokenRequestDTO, pub *rsa.PublicKey) (*dto.VerifyTokenResponseDTO, error) {
	token, err := jwt.Parse(req.JWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return pub, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	iss, _ := claims["iss"].(string)
	exp, _ := claims["exp"].(float64)

	return &dto.VerifyTokenResponseDTO{
		Valid:  true,
		Issuer: iss,
		Exp:    int64(exp),
	}, nil
}
