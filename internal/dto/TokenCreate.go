package dto

// TokenRequestDTO represents a request to create a JWT after verifying a TON wallet signature.
// swagger:model
type TokenRequestDTO struct {
	// Public key of the TON wallet in base64 format
	// required: true
	// example: dGVzdF9wdWJsaWNfa2V5X2RhdGE=
	PublicKey string `json:"publicKey" validate:"required,len=44" example:"dGVzdF9wdWJsaWNfa2V5X2RhdGE="`

	// Original message that was signed
	// required: true
	// example: TON OAuth challenge message
	Message string `json:"message" validate:"required" example:"TON OAuth challenge message"`

	// Signature of the message in base64 format
	// required: true
	// example: c2lnbmF0dXJlX2RhdGFfYmFzZTY0X2Zvcm1hdA==
	Signature string `json:"signature" validate:"required,len=88" example:"c2lnbmF0dXJlX2RhdGFfYmFzZTY0X2Zvcm1hdA=="`
}

// TokenResponseDTO represents the response containing the JWT token.
// swagger:model
type TokenResponseDTO struct {
	// JWT token issued after successful verification
	// required: true
	// example: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...
	JWT string `json:"jwt" validate:"required" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
