package usecase

import (
	"TON/internal/dto"
	"TON/pkg/logger"
	"TON/pkg/tonwallet"
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/xssnick/tonutils-go/address"
)

type VerifyUseCase interface {
	Verify(req dto.VerifyRequestDTO) (*dto.VerifyResponseDTO, error)
}

type VerifyUseCaseImpl struct {
	Issuer  string
	TTL     time.Duration
	log     logger.Logger
	checker *tonwallet.WalletChecker
}

func NewVerifyUseCase(issuer string, ttl time.Duration, log logger.Logger, checker *tonwallet.WalletChecker) VerifyUseCase {
	return &VerifyUseCaseImpl{
		Issuer:  issuer,
		TTL:     ttl,
		log:     log,
		checker: checker,
	}
}

func (u *VerifyUseCaseImpl) Verify(req dto.VerifyRequestDTO) (*dto.VerifyResponseDTO, error) {
	ctx := context.Background()
	u.log.Info(ctx, "Starting signature verification")

	if !ed25519.Verify(req.PublicKey, []byte(req.Message), req.Signature) {
		u.log.Error(ctx, "Invalid signature")
		return nil, errors.New("invalid signature")
	}

	parts := strings.Split(req.Message, ":")
	if len(parts) != 2 {
		u.log.Error(ctx, "Invalid message format")
		return nil, errors.New("invalid message format, expected 'issuer:timestamp'")
	}
	issuer, tsStr := parts[0], parts[1]

	if issuer != u.Issuer {
		u.log.Error(ctx, fmt.Sprintf("Invalid issuer: got %s, expected %s", issuer, u.Issuer))
		return nil, errors.New("invalid issuer")
	}

	ts, err := time.Parse(time.RFC3339, tsStr)
	if err != nil {
		u.log.Error(ctx, "Invalid timestamp format: "+err.Error())
		return nil, fmt.Errorf("invalid timestamp format: %w", err)
	}

	now := time.Now()
	if ts.After(now) {
		u.log.Error(ctx, "Timestamp is from the future")
		return nil, errors.New("timestamp is from the future")
	}
	if now.Sub(ts) > u.TTL {
		u.log.Error(ctx, "Message expired")
		return nil, errors.New("message expired")
	}

	hash := sha256.Sum256(req.PublicKey)
	tonAddr := address.NewAddress(0, 0, hash[:]).String()

	active, err := u.checker.IsWalletActive(ctx, tonAddr)
	if err != nil {
		u.log.Error(ctx, "Failed to check wallet activity: "+err.Error())
		return nil, fmt.Errorf("failed to check wallet activity: %w", err)
	}
	if !active {
		u.log.Error(ctx, "Wallet is not active: "+tonAddr)
		return nil, errors.New("wallet is not active")
	}

	u.log.Info(ctx, "Signature and wallet verification successful for wallet "+tonAddr)

	return &dto.VerifyResponseDTO{
		Valid:     true,
		Wallet:    tonAddr,
		Issuer:    issuer,
		ExpiresAt: ts.Add(u.TTL),
	}, nil
}
