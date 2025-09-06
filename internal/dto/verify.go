package dto

import (
	"time"
)

type VerifyRequestDTO struct {
	Message   string `json:"message" validate:"required" example:"TON OAuth challenge message"`
	Signature []byte `json:"signature" validate:"required,len=64" example:"c2lnbmF0dXJlX2RhdGFfYmFzZTY0X2Zvcm1hdA=="`
	PublicKey []byte `json:"publicKey" validate:"required,len=32" example:"dGVzdF9wdWJsaWNfa2V5X2RhdGE="`
}

type VerifyResponseDTO struct {
	Valid     bool      `json:"valid" example:"true"`
	Wallet    string    `json:"wallet" example:"EQC1234567890abcdef..."`
	Issuer    string    `json:"issuer" example:"TON OAuth Service"`
	Nonce     string    `json:"nonce" example:"nonce1234567890"`
	ExpiresAt time.Time `json:"expiresAt" example:"2025-09-07T00:00:00Z"`
}
