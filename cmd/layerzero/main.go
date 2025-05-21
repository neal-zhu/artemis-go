package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/neal-zhu/artemis-go/pkg/collector"
	"github.com/neal-zhu/artemis-go/pkg/core"
	"github.com/neal-zhu/artemis-go/pkg/executor"
)

// main is the entry point for the layerzero command line tool.
//
// It is responsible for setting up the environment and then executing the
// strategy.
func main() {
	appCtx, appCancel := context.WithCancel(context.Background()) // Renamed for clarity
	defer appCancel() // Ensures cancellation if main exits for other reasons (e.g., panic) or after engine.Stop() completes successfully.

	logCollector, err := collector.NewLogCollector(
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
		log.Fatalf("Failed to create log collector: %v", err) // Improved error message
	}

	engine := core.NewEngine()
	engine.AddCollector(logCollector)
	engine.AddStrategy(NewLayerZeroNewTokenStrategy())
	engine.SetExecutor(executor.NewDummyExecutor())

	// 启动引擎
	if err := engine.Start(appCtx); err != nil { // Pass the application context to the engine
		log.Fatalf("Failed to start engine: %v", err)
	}

	// 设置信号处理，以便优雅地关闭程序
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 等待中断信号
	<-sigChan
	log.Println("Shutting down...")

	// 显式取消应用程序上下文，通知所有监听此上下文的协程开始关闭
	appCancel() // Explicitly cancel the context to signal shutdown

	// 等待引擎完全停止 (Engine.Stop() 应该等待其所有内部协程完成)
	engine.Stop()

	log.Println("Shutdown complete")
}
