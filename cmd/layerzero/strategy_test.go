package main

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/neal-zhu/artemis-go/pkg/collector"
	"github.com/neal-zhu/artemis-go/pkg/core"
)

func TestStrategy(t *testing.T) {
	collector, err := collector.NewLogCollector(
		"http://bsc_bnb.rpc.cobo.one",
		[]common.Address{
			common.HexToAddress("0x1a44076050125825900e736c501f859c50fE728c"),
		},
		[][]common.Hash{
			{common.HexToHash("0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981")},
		},
		big.NewInt(48459929-1),
	)
	if err != nil {
		t.Fatal(err)
	}
	collector.Start(context.Background())
	event := <-collector.Events()
	rawLog := event.(core.LogEvent)
	strategy := NewLayerZeroNewTokenStrategy()
	strategy.Process(context.Background(), rawLog)
}
