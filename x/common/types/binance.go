package types

import (
	"fmt"
	"reflect"

	bsccommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BinanceAddress defines a standard ethereum address
type BinanceAddress bsccommon.Address

// NewBinanceAddress is a constructor function for BinanceAddress
func NewBinanceAddress(address string) BinanceAddress {
	return BinanceAddress(bsccommon.HexToAddress(address))
}

// Route should return the name of the module
func (bscAddr BinanceAddress) String() string {
	return bsccommon.Address(bscAddr).String()
}

// MarshalJSON marshals the ethereum address to JSON
func (bscAddr BinanceAddress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", bscAddr.String())), nil
}

// UnmarshalJSON unmarshal an ethereum address
func (bscAddr *BinanceAddress) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(bsccommon.Address{}), input, bscAddr[:])
}
