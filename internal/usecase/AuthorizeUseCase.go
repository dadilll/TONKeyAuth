package usecase

import (
	"TON/internal/dto"
	"TON/pkg/logger"
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"
)

type AuthorizeUseCase interface {
	Authorize(req dto.AuthorizeRequestDTO) (*dto.AuthorizeResponseDTO, error)
}

type AuthorizeUseCaseImpl struct {
	TTL    int
	logger logger.Logger
}

func NewAuthorizeUseCase(ttl int, log logger.Logger) AuthorizeUseCase {
	return &AuthorizeUseCaseImpl{
		TTL:    ttl,
		logger: log,
	}
}

func (u *AuthorizeUseCaseImpl) Authorize(req dto.AuthorizeRequestDTO) (*dto.AuthorizeResponseDTO, error) {
	ctx := context.Background()

	clientID, err := generateRandomString(16)
	if err != nil {
		u.logger.Error(ctx, "Failed to generate clientID: "+err.Error())
		return nil, err
	}

	nonce, err := generateRandomString(32)
	if err != nil {
		u.logger.Error(ctx, "Failed to generate nonce: "+err.Error())
		return nil, err
	}

	challenge := &dto.AuthorizeResponseDTO{
		ClientID:    clientID,
		RedirectURI: req.RedirectURI,
		Challenge:   nonce,
		ExpiresAt:   time.Now().Add(time.Duration(u.TTL) * time.Second),
	}

	u.logger.Info(ctx, "Generated anonymous challenge")

	return challenge, nil
}

func generateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
