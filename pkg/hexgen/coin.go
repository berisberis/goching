package hexgen

import (
	"crypto/rand"
	"math/big"
)

func CoinToss() (int64, error) {

	side, err := rand.Int(rand.Reader, big.NewInt(2))

	if err != nil {
		return 2, err
	}

	return side.Int64(), nil
}
