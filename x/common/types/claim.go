package types

// EthERC20ProphecyClaim contains data required to make an ProphecyClaim
type EthERC20ProphecyClaim struct {
	ClaimType      ClaimType       `json:"claim_type"`
	EthereumSender EthereumAddress `json:"ethereum_sender"`
	BinanceReceiver BinanceAddress   `json:"binance_receiver"`
	Symbol         string          `json:"symbol"`
	Amount         string          `json:"amount"`
	TxHash         string          `json:"tx_hash"`
}

// BscERC20ProphecyClaim contains data required to make an ProphecyClaim
type BscERC20ProphecyClaim struct {
	ChainName        string          `json:"chain_name"`
	ClaimType        ClaimType       `json:"claim_type"`
	BinanceSender    BinanceAddress  `json:"binance_sender"`
	EthereumReceiver EthereumAddress `json:"ethereum_receiver"`
	Symbol           string          `json:"symbol"`
	Amount           string          `json:"amount"`
	TxHash           string          `json:"tx_hash"`
}

// EthERC721ProphecyClaim contains data required to make an ProphecyClaim
type EthERC721ProphecyClaim struct {
	ClaimType       ClaimType       `json:"claim_type"`
	ChainName       string          `json:"chain_name"`
	EthereumSender  EthereumAddress `json:"ethereum_sender"`
	BinanceReceiver BinanceAddress  `json:"binance_receiver"`
	Symbol          string          `json:"symbol"`
	TokenId         string          `json:"token_id"`
	BaseURI         string          `json:"base_uri"`
	TokenURI        string          `json:"token_uri"`
	TxHash          string          `json:"tx_hash"`
}

// BscERC721ProphecyClaim contains data required to make an ProphecyClaim
type BscERC721ProphecyClaim struct {
	ClaimType        ClaimType       `json:"claim_type"`
	ChainName        string          `json:"chain_name"`
	BinanceSender    BinanceAddress  `json:"binance_sender"`
	EthereumReceiver EthereumAddress `json:"ethereum_receiver"`
	Symbol           string          `json:"symbol"`
	TokenId          string          `json:"token_id"`
	BaseURI          string          `json:"base_uri"`
	TokenURI         string          `json:"token_uri"`
	TxHash           string          `json:"tx_hash"`
}