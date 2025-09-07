package http

import (
	"TON/internal/config"
	"TON/internal/handler"
	"TON/internal/usecase"
	"TON/pkg/logger"
	"TON/pkg/tonwallet"
	"TON/pkg/validator"
	"crypto/rsa"
	"time"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, cfg *config.Config, log logger.Logger, privKey *rsa.PrivateKey, pubKey *rsa.PublicKey, tonwallet *tonwallet.WalletChecker) {

	authorizeUC := usecase.NewAuthorizeUseCase(120, log)
	verifyUC := usecase.NewVerifyUseCase(cfg.Issuer, 2*time.Minute, log, tonwallet)
	tokenUC := usecase.NewTokenUseCase(cfg.Issuer, 5*time.Minute, privKey)
	jwksUC := usecase.NewJWKSUseCase(cfg.KeyName, pubKey)
	verifyTokenUC := usecase.NewTokenVerifyUseCase()

	val := validator.NewCustomValidator()

	oauthHandler := handler.NewOauthHandler(
		log,
		val,
		authorizeUC,
		verifyUC,
		tokenUC,
		jwksUC,
		verifyTokenUC,
	)

	api := e.Group("/oauth")
	api.GET("/authorize", oauthHandler.AuthorizeHandler)
	api.POST("/verify", oauthHandler.VerifyHandler)
	api.POST("/token", oauthHandler.TokenHandler)
	api.GET("/jwks", oauthHandler.JWKSHandler)
	api.POST("/verify-token", oauthHandler.VerifyTokenHandler)
}
