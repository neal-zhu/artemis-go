package collector

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

type TransactionCollector struct {
	client    *ethclient.Client
	eventChan chan core.Event
	stopChan  chan struct{}
}

func NewTransactionCollector(url string) (*TransactionCollector, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	return &TransactionCollector{
		client:    client,
		eventChan: make(chan core.Event),
		stopChan:  make(chan struct{}),
	}, nil
}

func (c *TransactionCollector) Start(ctx context.Context) error {
	go c.collect(ctx)
	return nil
}

func (c *TransactionCollector) Stop() error {
	close(c.stopChan)
	return nil
}

func (c *TransactionCollector) Events() <-chan core.Event {
	return c.eventChan
}

func (c *TransactionCollector) collect(ctx context.Context) {
	headers := make(chan *types.Header)
	sub, err := c.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		fmt.Printf("Error subscribing to new headers: %v\n", err)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.stopChan:
			return
		case err := <-sub.Err():
			fmt.Printf("Error in header subscription: %v\n", err)
			return
		case header := <-headers:
			block, err := c.client.BlockByHash(ctx, header.Hash())
			if err != nil {
				fmt.Printf("Error fetching block: %v\n", err)
				continue
			}

			for _, tx := range block.Transactions() {
				from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
				if err != nil {
					fmt.Printf("Error getting transaction sender: %v\n", err)
					continue
				}

				c.eventChan <- core.TransactionEvent{
					TxHash: tx.Hash().Hex(),
					From:   from.Hex(),
					To:     tx.To().Hex(),
					Value:  tx.Value().String(),
				}
			}
		}
	}
}
