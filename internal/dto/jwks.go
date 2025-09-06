package dto

type JWKSResponseDTO struct {
	Keys []JWK `json:"keys" validate:"required,dive"`
}

type JWK struct {
	Kid string `json:"kid" validate:"required" example:"key1"`
	Kty string `json:"kty" validate:"required" example:"RSA"`
	Alg string `json:"alg" validate:"required" example:"RS256"`
	Use string `json:"use" validate:"required" example:"sig"`
	N   string `json:"n" validate:"required" example:"0vx7agoebGcQSuuPiLJXZptNnP9Z..."`
	E   string `json:"e" validate:"required" example:"AQAB"`
}
