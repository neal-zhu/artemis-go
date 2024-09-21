package core

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	ActionTypeExecuteTrade ActionType = iota
	ActionTypeUpdateState
	// Add more action types as needed
	ActionTypeSendTransaction
	ActionTypeFlashbots
	ActionTypeMEVShare
)

// ExecuteTradeAction represents an action to execute a trade
type ExecuteTradeAction struct {
	Asset     string
	Amount    string
	Price     string
	TradeType string
}

func (a ExecuteTradeAction) Type() ActionType {
	return ActionTypeExecuteTrade
}

// SendTransactionAction represents an action to send a transaction
type SendTransactionAction struct {
	From       common.Address
	To         common.Address
	Value      *big.Int
	Data       []byte
	GasLimit   uint64
	PrivateKey *ecdsa.PrivateKey
}

func (a SendTransactionAction) Type() ActionType {
	return ActionTypeSendTransaction
}

// FlashbotsAction represents an action to execute a Flashbots bundle
type FlashbotsAction struct {
	Transactions []*types.Transaction
	BlockNumber  *big.Int
	ChainID      *big.Int
	PrivateKey   *ecdsa.PrivateKey
}

func (a FlashbotsAction) Type() ActionType {
	return ActionTypeFlashbots
}

// MEVShareAction represents an action to execute a MEV-Share bundle
type MEVShareAction struct {
	Transactions [][]byte
	BlockNumber  *big.Int
	Hints        []string
}

func (a MEVShareAction) Type() ActionType {
	return ActionTypeMEVShare
}
