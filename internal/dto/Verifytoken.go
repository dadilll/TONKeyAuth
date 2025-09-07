package dto

// VerifyTokenRequestDTO represents a request to verify a JWT token.
// swagger:model
type VerifyTokenRequestDTO struct {
	// JWT token string to verify
	// required: true
	// example: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...
	JWT string `json:"jwt" validate:"required" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// VerifyTokenResponseDTO represents the response after verifying a JWT token.
// swagger:model
type VerifyTokenResponseDTO struct {
	// Indicates if the token is valid
	// required: true
	// example: true
	Valid bool `json:"valid" example:"true"`

	// Issuer of the JWT token
	// required: true
	// example: TON OAuth Service
	Issuer string `json:"issuer" example:"TON OAuth Service"`

	// Expiration timestamp of the token (Unix time)
	// required: true
	// example: 1751913600
	Exp int64 `json:"exp" example:"1751913600"`
}
