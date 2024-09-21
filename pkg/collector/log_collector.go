package collector

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

type LogCollector struct {
	client    *ethclient.Client
	eventChan chan core.Event
	stopChan  chan struct{}
	addresses []common.Address
	topics    [][]common.Hash
}

func NewLogCollector(url string, addresses []common.Address, topics [][]common.Hash) (*LogCollector, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	return &LogCollector{
		client:    client,
		eventChan: make(chan core.Event),
		stopChan:  make(chan struct{}),
		addresses: addresses,
		topics:    topics,
	}, nil
}

func (c *LogCollector) Start(ctx context.Context) error {
	go c.collect(ctx)
	return nil
}

func (c *LogCollector) Stop() error {
	close(c.stopChan)
	return nil
}

func (c *LogCollector) Events() <-chan core.Event {
	return c.eventChan
}

func (c *LogCollector) collect(ctx context.Context) {
	query := ethereum.FilterQuery{
		Addresses: c.addresses,
		Topics:    c.topics,
	}

	logs := make(chan types.Log)
	sub, err := c.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		fmt.Printf("Error subscribing to logs: %v\n", err)
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
			fmt.Printf("Error in log subscription: %v\n", err)
			return
		case log := <-logs:
			c.eventChan <- core.LogEvent{
				Address:     log.Address.Hex(),
				Topics:      log.Topics,
				Data:        log.Data,
				BlockNumber: log.BlockNumber,
				TxHash:      log.TxHash.Hex(),
				TxIndex:     log.TxIndex,
				BlockHash:   log.BlockHash.Hex(),
				Index:       log.Index,
			}
		}
	}
}
