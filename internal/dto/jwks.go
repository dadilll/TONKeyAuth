package dto

// JWKSResponseDTO represents the JSON Web Key Set response.
// swagger:model
type JWKSResponseDTO struct {
	// Array of JSON Web Keys
	// required: true
	Keys []JWK `json:"keys" validate:"required,dive"`
}

// JWK represents a single JSON Web Key.
// swagger:model
type JWK struct {
	// Key ID
	// required: true
	// example: key1
	Kid string `json:"kid" validate:"required" example:"key1"`

	// Key type
	// required: true
	// example: RSA
	Kty string `json:"kty" validate:"required" example:"RSA"`

	// Algorithm
	// required: true
	// example: RS256
	Alg string `json:"alg" validate:"required" example:"RS256"`

	// Intended use of the key
	// required: true
	// example: sig
	Use string `json:"use" validate:"required" example:"sig"`

	// Modulus for RSA public key
	// required: true
	// example: 0vx7agoebGcQSuuPiLJXZptNnP9Z...
	N string `json:"n" validate:"required" example:"0vx7agoebGcQSuuPiLJXZptNnP9Z..."`

	// Exponent for RSA public key
	// required: true
	// example: AQAB
	E string `json:"e" validate:"required" example:"AQAB"`
}
