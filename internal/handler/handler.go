package handler

import (
	"TON/internal/dto"
	"TON/internal/usecase"
	"TON/pkg/logger"
	"TON/pkg/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OauthHandler struct {
	AuthorizeUseCase   usecase.AuthorizeUseCase
	VerifyUseCase      usecase.VerifyUseCase
	TokenUseCase       usecase.TokenUseCase
	JWKSUseCase        usecase.JWKSUseCase
	TokenVerifyUseCase usecase.TokenVerifyUseCase
	logger             logger.Logger
	validator          *validator.CustomValidator
}

func NewOauthHandler(
	log logger.Logger,
	val *validator.CustomValidator,
	authorize usecase.AuthorizeUseCase,
	verify usecase.VerifyUseCase,
	token usecase.TokenUseCase,
	jwks usecase.JWKSUseCase,
	tokenVerify usecase.TokenVerifyUseCase,
) *OauthHandler {
	return &OauthHandler{
		logger:             log,
		validator:          val,
		AuthorizeUseCase:   authorize,
		VerifyUseCase:      verify,
		TokenUseCase:       token,
		JWKSUseCase:        jwks,
		TokenVerifyUseCase: tokenVerify,
	}
}

// AuthorizeHandler godoc
// @Summary Generate authorization challenge
// @Description Generate a one-time nonce for TON OAuth.
// @Tags auth
// @Accept json
// @Produce json
// @Param redirect_uri query string true "Redirect URI"
// @Param scope query string false "Scope"
// @Success 200 {object} dto.AuthorizeResponseDTO
// @Failure 400 {object} map[string]string "Bad request, invalid parameters"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /oauth/authorize [get]
func (h *OauthHandler) AuthorizeHandler(c echo.Context) error {
	req := dto.AuthorizeRequestDTO{
		RedirectURI: c.QueryParam("redirect_uri"),
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	resp, err := h.AuthorizeUseCase.Authorize(req)
	if err != nil {
		h.logger.Error(c.Request().Context(), "failed to generate challenge: "+err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate challenge"})
	}

	return c.JSON(http.StatusOK, resp)
}

// VerifyHandler godoc
// @Summary Verify TON wallet signature
// @Description Verify signed message from TON wallet using ed25519.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body object true "Verify request" schema="{\"type\":\"object\",\"properties\":{\"message\":{\"type\":\"string\",\"example\":\"TON OAuth challenge message\"},\"signature\":{\"type\":\"string\",\"example\":\"c2lnbmF0dXJlX2RhdGFfYmFzZTY0X2Zvcm1hdA==\"},\"publicKey\":{\"type\":\"string\",\"example\":\"dGVzdF9wdWJsaWNfa2V5X2RhdGE=\"}}}"
// @Success 200 {object} dto.VerifyResponseDTO
// @Failure 400 {object} map[string]string "Bad request, invalid body"
// @Failure 401 {object} map[string]string "Unauthorized, signature invalid"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /oauth/verify [post]
func (h *OauthHandler) VerifyHandler(c echo.Context) error {
	var req dto.VerifyRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	resp, err := h.VerifyUseCase.Verify(req)
	if err != nil {
		h.logger.Error(c.Request().Context(), "verification failed: "+err.Error())
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// TokenHandler godoc
// @Summary Create JWT token
// @Description Create JWT after successful verification of TON wallet.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.TokenRequestDTO true "Token request"
// @Success 200 {object} dto.TokenResponseDTO
// @Failure 400 {object} map[string]string "Bad request, invalid body"
// @Failure 401 {object} map[string]string "Unauthorized, verification failed"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /oauth/token [post]
func (h *OauthHandler) TokenHandler(c echo.Context) error {
	var req dto.TokenRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	resp, err := h.TokenUseCase.CreateToken(req)
	if err != nil {
		h.logger.Error(c.Request().Context(), "token creation failed: "+err.Error())
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// JWKSHandler godoc
// @Summary Get JSON Web Key Set (JWKS)
// @Description Get public keys to verify JWT tokens issued by TON OAuth.
// @Tags jwks
// @Produce json
// @Success 200 {object} dto.JWKSResponseDTO
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /oauth/jwks [get]
func (h *OauthHandler) JWKSHandler(c echo.Context) error {
	resp, err := h.JWKSUseCase.GetJWKS()
	if err != nil {
		h.logger.Error(c.Request().Context(), "failed to get JWKS: "+err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get JWKS"})
	}
	return c.JSON(http.StatusOK, resp)
}

// VerifyTokenHandler godoc
// @Summary Verify JWT token
// @Description Verify JWT issued by TON OAuth service.
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.VerifyTokenRequestDTO true "Token request"
// @Success 200 {object} dto.VerifyTokenResponseDTO
// @Failure 400 {object} map[string]string "Bad request, invalid body"
// @Failure 401 {object} map[string]string "Unauthorized, token invalid"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /oauth/verify-token [post]
func (h *OauthHandler) VerifyTokenHandler(c echo.Context) error {
	var req dto.VerifyTokenRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	pubKey := h.JWKSUseCase.GetPublicKey()
	resp, err := h.TokenVerifyUseCase.VerifyToken(req, pubKey)
	if err != nil {
		h.logger.Error(c.Request().Context(), "token verification failed: "+err.Error())
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
