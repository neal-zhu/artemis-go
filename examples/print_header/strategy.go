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
	if event.Type() != core.EventTypeLog {
		return nil, nil
	}

	headerEvent, ok := event.(core.LogEvent)
	if !ok {
		return nil, fmt.Errorf("invalid event type: expected LogEvent")
	}

	fmt.Printf("New block header received:\n")
	fmt.Printf("  Block Number: %d\n", headerEvent.BlockNumber)
	fmt.Printf("  Block Hash: %s\n", headerEvent.BlockHash)

	return nil, nil
}
