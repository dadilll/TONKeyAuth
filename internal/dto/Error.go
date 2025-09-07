package dto

// ErrorResponseDTO represents a standard error response
type ErrorResponseDTO struct {
	// Short error message
	// example: Validation failed
	Error string `json:"error"`
	// Optional detailed information about the error
	// example: "field 'redirect_uri' is required"
	Details string `json:"details,omitempty"`
}
