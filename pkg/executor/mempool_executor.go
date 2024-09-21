package executor

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

type MemPoolExecutor struct {
	client *ethclient.Client
}

func NewMemPoolExecutor(url string) (*MemPoolExecutor, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	return &MemPoolExecutor{client: client}, nil
}

func (e *MemPoolExecutor) Execute(ctx context.Context, action core.Action) error {
	sendTxAction, ok := action.(core.SendTransactionAction)
	if !ok {
		return fmt.Errorf("invalid action type for MemPoolExecutor: %T", action)
	}

	nonce, err := e.client.PendingNonceAt(ctx, sendTxAction.From)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}

	tx := types.NewTransaction(nonce, sendTxAction.To, sendTxAction.Value, sendTxAction.GasLimit, gasPrice, sendTxAction.Data)

	chainID, err := e.client.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), sendTxAction.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %v", err)
	}

	err = e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	return nil
}
