package types

// EthProphecyClaim contains data required to make an ProphecyClaim
type EthProphecyClaim struct {
	ClaimType      ClaimType       `json:"claim_type"`
	EthereumSender EthereumAddress `json:"ethereum_sender"`
	BinanceReceiver BinanceAddress   `json:"binance_receiver"`
	Symbol         string          `json:"symbol"`
	Amount         string          `json:"amount"`
	TxHash         string          `json:"tx_hash"`
}

// BscProphecyClaim contains data required to make an ProphecyClaim
type BscProphecyClaim struct {
	ChainName        string          `json:"chain_name"`
	ClaimType        ClaimType       `json:"claim_type"`
	BinanceSender     BinanceAddress   `json:"binance_sender"`
	EthereumReceiver EthereumAddress `json:"ethereum_receiver"`
	Symbol           string          `json:"symbol"`
	Amount           string          `json:"amount"`
	TxHash           string          `json:"tx_hash"`
}
