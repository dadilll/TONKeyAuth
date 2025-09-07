package main

import (
	"TON/internal/config"
	"TON/internal/transport/http"
	"TON/pkg/jwt"
	"TON/pkg/logger"
	"TON/pkg/tonwallet"
	"context"
)

const serviceName = "TonOauthService"

func main() {
	ctx := context.Background()
	Logger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, Logger)

	cfg := config.New()
	if cfg == nil {
		Logger.Error(ctx, "ERROR: config is nil")
		return
	}

	privateKey, err := jwt.LoadPrivateKey(cfg.PrivateKeyPath)
	if err != nil {
		Logger.Error(ctx, "Error loaded private key: "+err.Error())
	}

	publicKey, err := jwt.LoadPublicKey(cfg.PublicKeyPath)
	if err != nil {
		Logger.Error(ctx, "Error loaded private key: "+err.Error())
	}

	checker, err := tonwallet.NewWalletChecker(ctx)
	if err != nil {
		Logger.Error(ctx, "Error init TON WalletChecker: "+err.Error())
		return
	}

	e := http.New(Logger, cfg, privateKey, publicKey, checker)

	httpServer := http.Start(e, Logger, cfg.HTTPServerPort)

	http.WaitForShutdown(httpServer, Logger)
}
