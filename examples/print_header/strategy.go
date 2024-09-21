package main

import (
	"context"
	"fmt"

	"github.com/neal-zhu/artemis-go/pkg/core"
)

type SimpleHeaderPrintStrategy struct{}

func NewSimpleHeaderPrintStrategy() *SimpleHeaderPrintStrategy {
	return &SimpleHeaderPrintStrategy{}
}

func (s *SimpleHeaderPrintStrategy) Process(ctx context.Context, event core.Event) ([]core.Action, error) {
	if event.Type() != core.EventTypeBlockHeader {
		return nil, nil
	}

	headerEvent, ok := event.(core.BlockHeaderEvent)
	if !ok {
		return nil, fmt.Errorf("invalid event type: expected BlockHeaderEvent")
	}

	fmt.Printf("New block header received:\n")
	fmt.Printf("  Block Number: %d\n", headerEvent.BlockNumber)
	fmt.Printf("  Block Hash: %s\n", headerEvent.BlockHash)
	fmt.Printf("  Timestamp: %d\n", headerEvent.Timestamp)

	return nil, nil
}
