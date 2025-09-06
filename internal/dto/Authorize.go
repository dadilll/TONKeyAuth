package dto

import "time"

type AuthorizeRequestDTO struct {
	RedirectURI string `json:"redirect_uri" validate:"required,url" example:"https://example.com/callback"`
	Scope       string `json:"scope,omitempty" example:"read write"`
}

type AuthorizeResponseDTO struct {
	ClientID    string    `json:"client_id" example:"abc123def456"`
	RedirectURI string    `json:"redirect_uri" example:"https://example.com/callback"`
	Challenge   string    `json:"challenge" example:"nonce1234567890"`
	ExpiresAt   time.Time `json:"expiresAt" example:"2025-09-07T00:00:00Z"`
}
