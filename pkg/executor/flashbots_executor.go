package executor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/metachris/flashbotsrpc"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

type FlashbotsExecutor struct {
	fbClient   *flashbotsrpc.FlashbotsRPC
	chainID    *big.Int
	signingKey *ecdsa.PrivateKey
}

func NewFlashbotsExecutor(relayURL string, signingKey *ecdsa.PrivateKey, chainID *big.Int) *FlashbotsExecutor {
	fbClient := flashbotsrpc.NewFlashbotsRPC(relayURL)
	return &FlashbotsExecutor{
		fbClient:   fbClient,
		chainID:    chainID,
		signingKey: signingKey,
	}
}

func (e *FlashbotsExecutor) Execute(ctx context.Context, action core.Action) error {
	fbAction, ok := action.(core.FlashbotsAction)
	if !ok {
		return fmt.Errorf("invalid action type for FlashbotsExecutor: %T", action)
	}

	signedTxs := make([]*types.Transaction, len(fbAction.Transactions))
	for i, tx := range fbAction.Transactions {
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(e.chainID), e.signingKey)
		if err != nil {
			return fmt.Errorf("failed to sign transaction %d: %v", i, err)
		}
		signedTxs[i] = signedTx
	}

	txsHex := make([]string, len(signedTxs))
	for i, tx := range signedTxs {
		txBytes, err := tx.MarshalBinary()
		if err != nil {
			return fmt.Errorf("failed to marshal transaction %d: %v", i, err)
		}
		txsHex[i] = common.Bytes2Hex(txBytes)
	}

	blocknumber := fbAction.BlockNumber
	if blocknumber == nil {
		return fmt.Errorf("block number is required for Flashbots bundle")
	}

	simResult, err := e.fbClient.FlashbotsCallBundle(e.signingKey, flashbotsrpc.FlashbotsCallBundleParam{
		BlockNumber:      fmt.Sprintf("0x%x", blocknumber.Uint64()),
		Txs:              txsHex,
		StateBlockNumber: "latest",
	})
	if err != nil {
		return fmt.Errorf("Flashbots simulation failed: %v", err)
	}

	fmt.Printf("Simulation successful. Coinbase payment: %s wei\n", simResult.CoinbaseDiff)

	sendResult, err := e.fbClient.FlashbotsSendBundle(e.signingKey, flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         txsHex,
		BlockNumber: fmt.Sprintf("0x%x", blocknumber.Uint64()),
	})
	if err != nil {
		return fmt.Errorf("failed to send Flashbots bundle: %v", err)
	}

	fmt.Printf("Flashbots bundle sent. Bundle hash: %s\n", sendResult.BundleHash)
	return nil
}
