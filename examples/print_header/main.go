package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/neal-zhu/artemis-go/pkg/collector"
	"github.com/neal-zhu/artemis-go/pkg/core"
	"github.com/neal-zhu/artemis-go/pkg/executor"
)

func main() {
	// 定义命令行参数
	ethURL := flag.String("eth-url", "", "Ethereum node URL (required)")

	// 解析命令行参数
	flag.Parse()

	// 检查必要参数
	if *ethURL == "" {
		log.Fatal("Ethereum URL is required")
	}

	// 创建引擎
	engine := core.NewEngine()

	// 创建并添加区块头收集器
	headerCollector, err := collector.NewBlockHeaderCollector(*ethURL)
	if err != nil {
		log.Fatalf("Failed to create block header collector: %v", err)
	}
	engine.AddCollector(headerCollector)

	// 创建并添加简单的区块头打印策略
	headerPrintStrategy := NewSimpleHeaderPrintStrategy()
	engine.AddStrategy(headerPrintStrategy)

	// 创建并设置 DummyExecutor
	dummyExecutor := executor.NewDummyExecutor()
	engine.SetExecutor(dummyExecutor)

	// 创建一个可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动引擎
	if err := engine.Start(ctx); err != nil {
		log.Fatalf("Failed to start engine: %v", err)
	}

	// 设置信号处理，以便优雅地关闭程序
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 等待中断信号
	<-sigChan
	log.Println("Shutting down...")

	// 取消上下文，触发清理操作
	cancel()

	// 等待引擎完全停止
	engine.Stop()

	log.Println("Shutdown complete")
}
