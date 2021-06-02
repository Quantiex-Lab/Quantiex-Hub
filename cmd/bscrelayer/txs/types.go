package txs

import (
	"math/big"
)

// OracleClaim contains data required to make an OracleClaim
type OracleClaim struct {
	ProphecyID *big.Int
	Message    [32]byte
	Signature  []byte
}
