package collector

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

type BlockHeaderCollector struct {
	client    *ethclient.Client
	eventChan chan core.Event
	stopChan  chan struct{}
}

func NewBlockHeaderCollector(url string) (*BlockHeaderCollector, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	return &BlockHeaderCollector{
		client:    client,
		eventChan: make(chan core.Event),
		stopChan:  make(chan struct{}),
	}, nil
}

func (c *BlockHeaderCollector) Start(ctx context.Context) error {
	go c.collect(ctx)
	return nil
}

func (c *BlockHeaderCollector) Stop() error {
	close(c.stopChan)
	return nil
}

func (c *BlockHeaderCollector) Events() <-chan core.Event {
	return c.eventChan
}

func (c *BlockHeaderCollector) collect(ctx context.Context) {
	ticker := time.NewTicker(15 * time.Second) // 假设每15秒检查一次新区块
	defer ticker.Stop()

	var lastBlockNumber uint64

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.stopChan:
			return
		case <-ticker.C:
			header, err := c.client.HeaderByNumber(ctx, nil)
			if err != nil {
				fmt.Printf("Error fetching latest block header: %v\n", err)
				continue
			}

			if header.Number.Uint64() > lastBlockNumber {
				lastBlockNumber = header.Number.Uint64()
				c.eventChan <- core.BlockHeaderEvent{
					BlockNumber: header.Number.Uint64(),
					BlockHash:   header.Hash().Hex(),
					Timestamp:   int64(header.Time),
				}
			}
		}
	}
}
