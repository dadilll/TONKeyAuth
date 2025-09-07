package dto

import "time"

// VerifyRequestDTO represents a request to verify a signed message from a TON wallet.
// swagger:model
type VerifyRequestDTO struct {
	// Original message that was signed
	// required: true
	// example: TON OAuth challenge message
	Message string `json:"message" validate:"required" example:"TON OAuth challenge message"`

	// Signature of the message in base64 format
	// required: true
	// example: c2lnbmF0dXJlX2RhdGFfYmFzZTY0X2Zvcm1hdA==
	Signature []byte `json:"signature" validate:"required,len=64" example:"c2lnbmF0dXJlX2RhdGFfYmFzZTY0X2Zvcm1hdA=="`

	// Public key of the TON wallet in base64 format
	// required: true
	// example: dGVzdF9wdWJsaWNfa2V5X2RhdGE=
	PublicKey []byte `json:"publicKey" validate:"required,len=32" example:"dGVzdF9wdWJsaWNfa2V5X2RhdGE="`
}

// VerifyResponseDTO represents the response after verifying a TON wallet signature.
// swagger:model
type VerifyResponseDTO struct {
	// Indicates if the signature is valid
	// required: true
	// example: true
	Valid bool `json:"valid" example:"true"`

	// Wallet address of the signer
	// required: true
	// example: EQC1234567890abcdef...
	Wallet string `json:"wallet" example:"EQC1234567890abcdef..."`

	// Issuer of the verification
	// required: true
	// example: TON OAuth Service
	Issuer string `json:"issuer" example:"TON OAuth Service"`

	// Nonce used in the signed message
	// required: true
	// example: nonce1234567890
	Nonce string `json:"nonce" example:"nonce1234567890"`

	// Expiration timestamp of the verification
	// required: true
	// example: 2025-09-07T00:00:00Z
	ExpiresAt time.Time `json:"expiresAt" example:"2025-09-07T00:00:00Z"`
}
