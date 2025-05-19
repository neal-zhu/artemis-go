package main

import "github.com/neal-zhu/artemis-go/pkg/core"

package main

import (
	"context"
	"fmt"

	"github.com/neal-zhu/artemis-go/pkg/core"
)

type LayerZeroNewTokenStrategy struct{}

func NewLayerZeroNewTokenStrategy() *LayerZeroNewTokenStrategy {
	return &LayerZeroNewTokenStrategy{}
}

func (s *LayerZeroNewTokenStrategy) Process(ctx context.Context, event core.Event) ([]core.Action, error) {
	if event.Type() != core.EventTypeLog {
		return nil, nil
	}

	logEvent, ok := event.(core.LogEvent)
	if !ok {
		return nil, fmt.Errorf("invalid event type: expected LogEvent")
	}

	// check


	return nil, nil
}
