package usecase

import (
	"TON/internal/dto"
	"TON/pkg/logger"
	"context"
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type VerifyUseCase interface {
	Verify(req dto.VerifyRequestDTO) (*dto.VerifyResponseDTO, error)
}

type VerifyUseCaseImpl struct {
	Issuer string
	TTL    time.Duration
	log    logger.Logger
	apiURL string
	apiKey string
	client *http.Client
}

func NewVerifyUseCase(issuer string, ttl time.Duration, log logger.Logger, apiURL, apiKey string) VerifyUseCase {
	return &VerifyUseCaseImpl{
		Issuer: issuer,
		TTL:    ttl,
		log:    log,
		apiURL: apiURL,
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
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

	walletAddr, err := u.getWalletFromTonAPI(req.PublicKey)
	if err != nil {
		u.log.Error(ctx, "Failed to resolve wallet from TonAPI: "+err.Error())
		return nil, fmt.Errorf("failed to resolve wallet: %w", err)
	}

	active, err := u.isWalletActive(walletAddr)
	if err != nil {
		u.log.Error(ctx, "Failed to check wallet activity: "+err.Error())
		return nil, fmt.Errorf("failed to check wallet activity: %w", err)
	}
	if !active {
		u.log.Error(ctx, "Wallet is not active: "+walletAddr)
		return nil, errors.New("wallet is not active")
	}

	u.log.Info(ctx, "Signature and wallet verification successful for wallet "+walletAddr)
	return &dto.VerifyResponseDTO{
		Valid:     true,
		Wallet:    walletAddr,
		Issuer:    issuer,
		ExpiresAt: ts.Add(u.TTL),
	}, nil
}

func (u *VerifyUseCaseImpl) getWalletFromTonAPI(publicKey []byte) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/wallets?public_key=%x", u.apiURL, publicKey), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+u.apiKey)

	resp, err := u.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("TonAPI returned status %d", resp.StatusCode)
	}

	var data struct {
		Wallets []struct {
			Address string `json:"address"`
		} `json:"wallets"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if len(data.Wallets) == 0 {
		return "", errors.New("wallet not found")
	}

	return data.Wallets[0].Address, nil
}

func (u *VerifyUseCaseImpl) isWalletActive(addr string) (bool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/accounts/%s", u.apiURL, addr), nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+u.apiKey)

	resp, err := u.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("TonAPI returned status %d", resp.StatusCode)
	}

	var data struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return false, err
	}

	return data.Status == "active", nil
}
