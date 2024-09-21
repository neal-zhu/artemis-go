package core

import "github.com/ethereum/go-ethereum/common"

const (
	EventTypeBlockHeader EventType = iota
	EventTypeTransaction
	// Add more event types as needed
	EventTypeLog
)

// BlockHeaderEvent represents a new block header event
type BlockHeaderEvent struct {
	BlockNumber uint64
	BlockHash   string
	Timestamp   int64
}

func (e BlockHeaderEvent) Type() EventType {
	return EventTypeBlockHeader
}

// TransactionEvent represents a new transaction event
type TransactionEvent struct {
	TxHash string
	From   string
	To     string
	Value  string
}

func (e TransactionEvent) Type() EventType {
	return EventTypeTransaction
}

// LogEvent represents a new log event
type LogEvent struct {
	Address     string
	Topics      []common.Hash
	Data        []byte
	BlockNumber uint64
	TxHash      string
	TxIndex     uint
	BlockHash   string
	Index       uint
}

func (e LogEvent) Type() EventType {
	return EventTypeLog
}
