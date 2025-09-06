package dto

type TokenRequestDTO struct {
	PublicKey string `json:"publicKey" validate:"required,len=44" example:"dGVzdF9wdWJsaWNfa2V5X2RhdGE="`
	Message   string `json:"message" validate:"required" example:"TON OAuth challenge message"`
	Signature string `json:"signature" validate:"required,len=88" example:"c2lnbmF0dXJlX2RhdGFfYmFzZTY0X2Zvcm1hdA=="`
}
type TokenResponseDTO struct {
	JWT string `json:"jwt" validate:"required" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
