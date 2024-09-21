package core

import (
	"context"
)

// EventType represents the type of an event
type EventType int

// ActionType represents the type of an action
type ActionType int

// Event represents a generic event in the system
type Event interface {
	Type() EventType
}

// Action represents a generic action to be executed
type Action interface {
	Type() ActionType
}

// Collector collects external events and converts them to internal events
type Collector interface {
	Start(ctx context.Context) error
	Stop() error
	Events() <-chan Event
}

// Strategy processes events and generates actions
type Strategy interface {
	Process(ctx context.Context, event Event) ([]Action, error)
}

// Executor executes actions
type Executor interface {
	Execute(ctx context.Context, action Action) error
}
