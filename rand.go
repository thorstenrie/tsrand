package tsrand

import (
	"math/rand"

	"github.com/thorstenrie/tserr"
)

func CryptoRand() (*rand.Rand, error) {
	return New(cryptoSource())
}

func NewDeterministicRand() (*rand.Rand, error) {
	return New(newDeterministicSource())
}

// Seed and Read not secure for concurrency for sources from NewSource()!
func New(src Source) (*rand.Rand, error) {
	if src.Assert(); src.Err() != nil {
		return nil, tserr.NotAvailable(&tserr.NotAvailableArgs{S: "rand source", Err: src.Err()})
	}
	return rand.New(src), nil
}
