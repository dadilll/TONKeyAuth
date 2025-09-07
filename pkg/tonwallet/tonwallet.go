package tonwallet

import (
	"context"
	"fmt"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

type WalletChecker struct {
	api *ton.APIClient
}

func NewWalletChecker(ctx context.Context) (*WalletChecker, error) {
	client := liteclient.NewConnectionPool()
	err := client.AddConnectionsFromConfigUrl(ctx, "https://ton.org/global.config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to TON network: %w", err)
	}

	api := ton.NewAPIClient(client)
	return &WalletChecker{api: api}, nil
}

func (wc *WalletChecker) IsWalletActive(ctx context.Context, rawAddr string) (bool, error) {
	addr, err := address.ParseAddr(rawAddr)
	if err != nil {
		return false, fmt.Errorf("invalid wallet address: %w", err)
	}

	mcInfo, err := wc.api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get masterchain info: %w", err)
	}

	acc, err := wc.api.GetAccount(ctx, mcInfo, addr)
	if err != nil {
		return false, fmt.Errorf("failed to get account state: %w", err)
	}

	return acc.IsActive, nil
}
