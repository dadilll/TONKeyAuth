package dto

type VerifyTokenRequestDTO struct {
	JWT string `json:"jwt" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type VerifyTokenResponseDTO struct {
	Valid  bool   `json:"valid" example:"true"`
	Issuer string `json:"issuer" example:"TON OAuth Service"`
	Exp    int64  `json:"exp" example:"1751913600"`
}
