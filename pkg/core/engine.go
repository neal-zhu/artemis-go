package core

import (
	"context"
	"fmt"
	"sync"
)

type Engine struct {
	collectors []Collector
	strategies []Strategy
	executor   Executor
	wg         sync.WaitGroup
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) AddCollector(collector Collector) {
	e.collectors = append(e.collectors, collector)
}

func (e *Engine) AddStrategy(strategy Strategy) {
	e.strategies = append(e.strategies, strategy)
}

func (e *Engine) SetExecutor(executor Executor) {
	e.executor = executor
}

func (e *Engine) Start(ctx context.Context) error {
	if e.executor == nil {
		return fmt.Errorf("executor is not set")
	}

	for _, collector := range e.collectors {
		e.wg.Add(1)
		go func(c Collector) {
			defer e.wg.Done()
			if err := c.Start(ctx); err != nil {
				fmt.Printf("Error starting collector: %v\n", err)
				return
			}
			for {
				select {
				case event := <-c.Events():
					e.processEvent(ctx, event)
				case <-ctx.Done():
					return
				}
			}
		}(collector)
	}

	return nil
}

func (e *Engine) Stop() error {
	for _, collector := range e.collectors {
		if err := collector.Stop(); err != nil {
			fmt.Printf("Error stopping collector: %v\n", err)
		}
	}
	e.wg.Wait()
	return nil
}

func (e *Engine) processEvent(ctx context.Context, event Event) {
	for _, strategy := range e.strategies {
		actions, err := strategy.Process(ctx, event)
		if err != nil {
			fmt.Printf("Error processing event: %v\n", err)
			continue
		}
		for _, action := range actions {
			if err := e.executor.Execute(ctx, action); err != nil {
				fmt.Printf("Error executing action: %v\n", err)
			}
		}
	}
}
