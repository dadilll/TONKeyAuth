package dto

import "time"

// AuthorizeRequestDTO represents the request for generating an authorization challenge.
// swagger:model
type AuthorizeRequestDTO struct {
	// Redirect URI where the user will be redirected after successful authorization
	// required: true
	// example: https://example.com/callback
	RedirectURI string `json:"redirect_uri" validate:"required,url"`

	// Optional scope of the access request
	// example: read write
	Scope string `json:"scope,omitempty"`
}

// AuthorizeResponseDTO represents the response containing the authorization challenge.
// swagger:model
type AuthorizeResponseDTO struct {
	// Client ID of the application requesting authorization
	// example: abc123def456
	ClientID string `json:"client_id"`

	// Redirect URI to which the user will be sent after authorization
	// example: https://example.com/callback
	RedirectURI string `json:"redirect_uri"`

	// Unique challenge (nonce) to be signed by the user's TON wallet
	// example: nonce1234567890
	Challenge string `json:"challenge"`

	// Expiration time of the challenge
	// example: 2025-09-07T00:00:00Z
	ExpiresAt time.Time `json:"expiresAt"`
}
