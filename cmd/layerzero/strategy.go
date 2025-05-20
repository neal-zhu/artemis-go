package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/neal-zhu/artemis-go/generated/layerzero"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

type LayerZeroNewOFTAction struct {
	token               common.Address
	createAtBlockNumber uint64
}

func (a LayerZeroNewOFTAction) Type() core.ActionType {
	return core.ActionTypeSendTransaction
}

type LayerZeroNewTokenStrategy struct {
	l0 *layerzero.Layerzero
}

func NewLayerZeroNewTokenStrategy() *LayerZeroNewTokenStrategy {
	l0, _ := layerzero.NewLayerzero(
		common.Address{},
		nil,
	)
	return &LayerZeroNewTokenStrategy{
		l0: l0,
	}
}

func (s *LayerZeroNewTokenStrategy) Process(ctx context.Context, event core.Event) ([]core.Action, error) {
	if event.Type() != core.EventTypeLog {
		return nil, nil
	}

	logEvent, ok := event.(core.LogEvent)
	if !ok {
		return nil, fmt.Errorf("invalid event type: expected LogEvent")
	}

	delegateSet, err := s.l0.ParseDelegateSet(logEvent.Log)
	if err != nil {
		return nil, err
	}

	return []core.Action{LayerZeroNewOFTAction{
		token:               delegateSet.Sender,
		createAtBlockNumber: logEvent.BlockNumber,
	}}, nil
}
