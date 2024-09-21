package executor

import (
	"context"
	"fmt"

	"github.com/neal-zhu/artemis-go/pkg/core"
)

// DummyExecutor 是一个简单的执行器，只输出日志
type DummyExecutor struct{}

func NewDummyExecutor() *DummyExecutor {
	return &DummyExecutor{}
}

func (e *DummyExecutor) Execute(ctx context.Context, action core.Action) error {
	fmt.Printf("DummyExecutor: Received action of type %v\n", action.Type())
	return nil
}
