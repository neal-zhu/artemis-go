package bsc_sandwich

// Config holds the configuration for the BSC sandwich strategy.
type Config struct {
	NodeURL           string                `json:"nodeURL" yaml:"nodeURL"`                       // BSC JSON-RPC node URL (e.g., for HTTP provider)
	WebSocketURL      string                `json:"webSocketURL" yaml:"webSocketURL"`             // BSC WebSocket URL (for mempool subscription)
	PrivateKey        string                `json:"privateKey" yaml:"privateKey"`                 // Private key of the wallet executing strategy transactions
	SandwichContract  ContractConfig        `json:"sandwichContract" yaml:"sandwichContract"`   // Configuration for our deployed SandwichSwap contract
	PancakeSwapRouter string                `json:"pancakeSwapRouter" yaml:"pancakeSwapRouter"` // Address of the PancakeSwap Router V2
	WBNBAddress       string                `json:"wbnbAddress" yaml:"wbnbAddress"`               // Address of WBNB on BSC
	BloXrouteConfig   BloXrouteAPIConfig    `json:"bloXrouteConfig" yaml:"bloXrouteConfig"`       // Configuration for bloXroute API
	MonitoredSettings MonitoredUserSettings `json:"monitoredSettings" yaml:"monitoredSettings"` // Settings for tokens and users to monitor
	StrategyParams    StrategyParameters    `json:"strategyParams" yaml:"strategyParams"`       // General strategy parameters
}

// ContractConfig contains details about our deployed smart contract.
type ContractConfig struct {
	Address         string `json:"address" yaml:"address"`                 // Address of the deployed SandwichSwap contract
	GasLimit        uint64 `json:"gasLimit" yaml:"gasLimit"`               // Default gas limit for transactions to our contract
	MinProfitMargin uint64 `json:"minProfitMargin" yaml:"minProfitMargin"` // Minimum profit margin in BNB (wei) to execute a sandwich
}

// BloXrouteAPIConfig holds configuration for the bloXroute MEV relay.
type BloXrouteAPIConfig struct {
	AuthHeader string `json:"authHeader" yaml:"authHeader"` // Authorization header value for bloXroute API
	Endpoint   string `json:"endpoint" yaml:"endpoint"`     // bloXroute API endpoint (e.g., "https://api.blxrbdn.com")
}

// MonitoredUserSettings defines which tokens and user activities to monitor.
type MonitoredUserSettings struct {
	TargetTokens      []TargetToken      `json:"targetTokens" yaml:"targetTokens"`           // List of tokens to potentially sandwich
	MonitoredWallets  []MonitoredWallet  `json:"monitoredWallets" yaml:"monitoredWallets"`   // Specific wallets to monitor for target transactions
	TransactionFilter TransactionMatcher `json:"transactionFilter" yaml:"transactionFilter"` // General filter for identifying target transactions
}

// TargetToken specifies a token and its sandwiching parameters.
type TargetToken struct {
	Address          string `json:"address" yaml:"address"`                     // Address of the target ERC20 token
	MinSwapAmount    uint64 `json:"minSwapAmount" yaml:"minSwapAmount"`         // Minimum amount of this token in a swap to consider sandwiching (in token's smallest unit)
	MaxSandwichAmount uint64 `json:"maxSandwichAmount" yaml:"maxSandwichAmount"` // Maximum amount of this token our strategy will use for the front-run part of the sandwich
}

// MonitoredWallet specifies a wallet and custom settings if any.
type MonitoredWallet struct {
	Address string `json:"address" yaml:"address"` // Wallet address to monitor
	// Potentially add specific override settings for this wallet later
}

// TransactionMatcher defines criteria for identifying interesting mempool transactions.
// This is kept flexible as per user's request ("具体逻辑先留白").
type TransactionMatcher struct {
	ToAddress        string   `json:"toAddress" yaml:"toAddress"`               // Transaction must be to this address (e.g., PancakeSwap Router)
	FunctionSignatures []string `json:"functionSignatures" yaml:"functionSignatures"` // List of 4-byte function signatures (e.g., "0x7ff36ab5" for swapExactETHForTokens)
	MinBNBValue      string   `json:"minBNBValue" yaml:"minBNBValue"`           // Minimum BNB value in the transaction (as string to handle large numbers, e.g., "0.1")
}

// StrategyParameters holds other operational parameters for the strategy.
type StrategyParameters struct {
	FrontrunGasPriceOffset string `json:"frontrunGasPriceOffset" yaml:"frontrunGasPriceOffset"` // Gas price offset (Gwei) for front-run tx relative to target tx. Can be positive or negative.
	BackrunGasPriceOffset  string `json:"backrunGasPriceOffset" yaml:"backrunGasPriceOffset"`   // Gas price offset (Gwei) for back-run tx relative to target tx.
	BundleAttempts         int    `json:"bundleAttempts" yaml:"bundleAttempts"`                 // How many blocks to attempt submitting the bundle for.
	SlippageToleranceBPS   uint64 `json:"slippageToleranceBPS" yaml:"slippageToleranceBPS"`     // Slippage tolerance in Basis Points (e.g., 50 for 0.5%) for our swaps.
}
