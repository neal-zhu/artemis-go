package collector

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

const (
	defaultPollingInterval    = 10 * time.Second
	defaultLookbackBlocks     = 100 // Number of blocks to look back on first poll if no specific start block is given
	defaultBlockConfirmations = 3   // Number of confirmations before considering a block final for polling
	defaultQueryRange         = 1000
)

type LogCollector struct {
	client             *ethclient.Client
	eventChan          chan core.Event
	stopChan           chan struct{}
	addresses          []common.Address
	topics             [][]common.Hash
	rpcUrl             string
	pollingBlockNumber *big.Int      // Last block number successfully polled + 1
	pollingInterval    time.Duration // How often to poll for new logs
	lookbackBlocks     int64         // For the initial poll, how many blocks to look back from the latest
	blockConfirmations uint64        // How many blocks behind the head to query, to avoid reorgs
}

func NewLogCollector(rpcURL string, addresses []common.Address, topics [][]common.Hash, fromBlockNumber *big.Int) (*LogCollector, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	return &LogCollector{
		client:             client,
		eventChan:          make(chan core.Event),
		stopChan:           make(chan struct{}),
		addresses:          addresses,
		topics:             topics,
		rpcUrl:             rpcURL,
		pollingInterval:    defaultPollingInterval,
		lookbackBlocks:     defaultLookbackBlocks,
		blockConfirmations: defaultBlockConfirmations,
		pollingBlockNumber: fromBlockNumber,
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
	// Ensure eventChan is initialized (it should be by NewLogCollector, but good practice)
	if c.eventChan == nil {
		// This case should ideally not happen if NewLogCollector is always used.
		c.eventChan = make(chan core.Event) // Consider a sensible default or panic.
		fmt.Println("LogCollector: Warning - eventChan was nil and re-initialized in collect. This might indicate an issue.") // TODO: use logger
	}

	ticker := time.NewTicker(c.pollingInterval)

	// Defer the closing of eventChan and stopping of the ticker to ensure cleanup on exit.
	defer func() {
		ticker.Stop()
		close(c.eventChan) // Close the event channel so consumers know there are no more events.
		fmt.Println("LogCollector: eventChan closed and ticker stopped.") // TODO: use a logger
	}()

	// Initialize pollingBlockNumber on the first run if it was not provided (i.e., nil).
	if c.pollingBlockNumber == nil {
		latestBlock, err := c.client.BlockNumber(ctx)
		if err != nil {
			fmt.Printf("LogCollector: Failed to get initial latest block number for %s: %v\n", c.rpcUrl, err) // TODO: use a logger
			return                                                                                            // Cannot proceed without a starting block determination.
		}
		initialFromBlock := int64(latestBlock) - c.lookbackBlocks
		if initialFromBlock < 0 {
			initialFromBlock = 0
		}
		c.pollingBlockNumber = big.NewInt(initialFromBlock)
		fmt.Printf("LogCollector: Initial pollingBlockNumber for %s set to %s (using lookback %d blocks from latest %d)\n", c.rpcUrl, c.pollingBlockNumber.String(), c.lookbackBlocks, latestBlock) // TODO: use a logger
	}

	fmt.Printf("LogCollector: Starting polling for %s, from initial block %s\n", c.rpcUrl, c.pollingBlockNumber.String()) // TODO: use a logger

	for {
		select {
		case <-ctx.Done():
			fmt.Println("LogCollector: Context done, stopping polling.") // TODO: use a logger
			return
		case <-c.stopChan:
			fmt.Println("LogCollector: Stop channel closed, stopping polling.") // TODO: use a logger
			return
		case <-ticker.C:
			latestBlockNumberFetched, err := c.client.BlockNumber(ctx)
			if err != nil {
				fmt.Printf("LogCollector: Failed to get latest block number for %s: %v\n", c.rpcUrl, err) // TODO: use a logger
				continue                                                                                  // Skip this tick, try again on the next one.
			}

			targetOverallToBlock := latestBlockNumberFetched - c.blockConfirmations

			if c.pollingBlockNumber.Uint64() > targetOverallToBlock {
				// fmt.Printf("LogCollector: Already caught up for %s. Current: %s, Target: %d\n", c.rpcUrl, c.pollingBlockNumber.String(), targetOverallToBlock) // TODO: use a logger with debug level
				continue // Already caught up or ahead.
			}

			// Inner loop to process blocks in chunks until caught up to targetOverallToBlock for this tick.
			for c.pollingBlockNumber.Uint64() <= targetOverallToBlock {
				// Check for shutdown signals before processing the next chunk.
				select {
				case <-ctx.Done():
					fmt.Println("LogCollector: Context done during chunk processing, stopping.") // TODO: use a logger
					return
				case <-c.stopChan:
					fmt.Println("LogCollector: Stop channel closed during chunk processing, stopping.") // TODO: use a logger
					return
				default:
					// Continue to process this chunk.
				}

				currentChunkFromBlock := new(big.Int).Set(c.pollingBlockNumber)
				currentChunkToBlockTarget := c.pollingBlockNumber.Uint64() + defaultQueryRange - 1

				var currentChunkToBlockUint64 uint64
				if currentChunkToBlockTarget > targetOverallToBlock {
					currentChunkToBlockUint64 = targetOverallToBlock
				} else {
					currentChunkToBlockUint64 = currentChunkToBlockTarget
				}

				// This should not happen if logic is correct, but as a safeguard:
				if currentChunkFromBlock.Uint64() > currentChunkToBlockUint64 {
					// fmt.Printf("LogCollector: FromBlock %s is greater than calculated ToBlock %d for %s. Breaking chunk loop.\n", currentChunkFromBlock.String(), currentChunkToBlockUint64, c.rpcUrl) // TODO: use a logger with debug level
					break // Break from this inner chunk-processing loop for this tick.
				}

				currentChunkToBlockBigInt := new(big.Int).SetUint64(currentChunkToBlockUint64)

				query := ethereum.FilterQuery{
					FromBlock: currentChunkFromBlock,
					ToBlock:   currentChunkToBlockBigInt,
					Addresses: c.addresses,
					Topics:    c.topics,
				}

				// fmt.Printf("LogCollector: Querying %s logs from %s to %s (Overall target: %d)\n", c.rpcUrl, query.FromBlock.String(), query.ToBlock.String(), targetOverallToBlock) // TODO: use a logger with debug level

				logs, err := c.client.FilterLogs(ctx, query)
				if err != nil {
					fmt.Printf("LogCollector: Error filtering logs for %s, range %s-%s: %v\n", c.rpcUrl, query.FromBlock.String(), query.ToBlock.String(), err) // TODO: use a logger
					// Break inner loop on error; will retry from current c.pollingBlockNumber in the next tick.
					break
				}

				// fmt.Printf("LogCollector: Found %d logs for %s between %s and %s\n", len(logs), c.rpcUrl, query.FromBlock.String(), query.ToBlock.String()) // TODO: use a logger with debug level
				for _, logEntry := range logs {
					select {
					case c.eventChan <- core.LogEvent{Log: logEntry}:
					case <-ctx.Done():
						fmt.Println("LogCollector: Context done during event send, stopping.") // TODO: use a logger
						return
					case <-c.stopChan:
						fmt.Println("LogCollector: Stop channel closed during event send, stopping.") // TODO: use a logger
						return
					}
				}

				c.pollingBlockNumber.SetUint64(currentChunkToBlockUint64 + 1)

				if currentChunkToBlockUint64 >= targetOverallToBlock {
					break // Reached overall target for this tick.
				}
			} // End of inner chunk-processing loop.
		}
	}
}
