package common

import (
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// enumer -type=Chain -linecomment -json=true -text=true -sql=true
type Chain uint64

const (
	ChainBase Chain = iota + 1 // base
)

// CrawlResult is result of the crawl
type CrawlResult struct {
	Trades []Tradelog `json:"trades"`
}

type Routes struct {
	LogIndex        uint             `json:"log_index"`
	TokenInAddress  ethereum.Address `json:"token_in_address"`
	TokenInAmount   float64          `json:"token_in_amount"`
	TokenInUsdtRate float64          `json:"token_in_usdt_rate"`
	AmountIn        *big.Int         `json:"-"`

	TokenOutAddress  ethereum.Address `json:"token_out_address"`
	TokenOutAmount   float64          `json:"token_out_amount"`
	TokenOutUsdtRate float64          `json:"token_out_usdt_rate"`
	AmountOut        *big.Int         `json:"-"`
	Exchange         ethereum.Address `json:"exchange"`
}

type Tradelog struct {
	BlockTimestamp time.Time        `json:"timestamp"`
	BlockNumber    uint64           `json:"block_number"`
	TxIndex        uint             `json:"tx_index"`
	TxHash         ethereum.Hash    `json:"tx_hash"`
	Sender         ethereum.Address `json:"sender"`

	TokenInAddress  ethereum.Address `json:"token_in_address"`
	TokenInAmount   float64          `json:"token_in_amount"`
	AmountIn        *big.Int
	TokenInUsdtRate float64 `json:"token_in_usdt_rate"`

	TokenOutAddress  ethereum.Address `json:"token_out_address"`
	TokenOutAmount   float64          `json:"token_out_amount"`
	AmountOut        *big.Int
	TokenOutUsdtRate float64 `json:"token_out_usdt_rate"`
	Routes           []byte  `json:"routes"`

	NativeTokenUsdtRate float64 `json:"native_token_usdt_rate"`
}

type Transferlog struct {
	BlockTimestamp time.Time     `json:"timestamp"`
	BlockNumber    uint64        `json:"block_number"`
	TxIndex        uint          `json:"tx_index"`
	TxHash         ethereum.Hash `json:"tx_hash"`

	FromAddress ethereum.Address `json:"from_address"`
	ToAddress   ethereum.Address `json:"to_address"`

	TokenAddress ethereum.Address `json:"token_in_address"`
	TokenAmount  float64          `json:"token_in_amount"`
	IsCexIn      bool             `json:"is_cex_in"`
	Exchange     ethereum.Address `json:"exchange"`
}

type SwapLogData struct {
	AmountIn         *big.Int
	AmountOut        *big.Int
	Address          ethereum.Address
	RecipientAddress ethereum.Address
	TokenIn          ethereum.Address
	TokenOut         ethereum.Address
}

type TransferLogData struct {
	From    ethereum.Address
	To      ethereum.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
	IsCexIn bool
}

type TokenContractInfo struct {
	Decimal uint8
	Symbol  string
	Name    string
}

// human-friendly
type TokenBalance struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type EvmConfig struct {
	Enable              bool
	WsKey               string
	FromBlock           int
	MaxWorkers          int
	MaxBlocksForWorker  int
	WaitForConfirmation int
}
