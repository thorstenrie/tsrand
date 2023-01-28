// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library packages and tserr
import (
	"math/rand" // math/rand
	"time"      // time

	"github.com/thorstenrie/tserr" // tserr
)

// defaultSeed is the default seed for the pseudo-random number generator
const (
	defaultSeed int64 = 1
)

// NewCryptoRand returns a new instance of rand.Rand which provides a cryptographically secure
// random number generator based on crypto/rand. It is safe for concurrent use by multiple goroutines.
// If it is not available on the platform, NewCryptoRand() returns an error and *rand.Rand is nil.
func NewCryptoRand() (*rand.Rand, error) {
	return New(cryptoSource())
}

// NewPseudoRandomRand returns a new instance of rand.Rand which provides a pseudo-
// random number generator based on math/rand. It is safe for concurrent use by multiple goroutines.
// The output might be easily predictable and is unsuitable for security-sensitive services.
func NewPseudoRandomRand() (*rand.Rand, error) {
	return New(newDeterministicSource(time.Now().UnixNano()))
}

// NewDeterministicRand returns a new instance of rand.Rand which provides a deterministic pseudo-
// random number generator based on math/rand. It is safe for concurrent use by multiple goroutines.
// It is initialized with defaultSeed and returns a deterministic random sequence. The output is
// easily predictable and is unsuitable for security-sensitive services.
func NewDeterministicRand() (*rand.Rand, error) {
	return New(newDeterministicSource(defaultSeed))
}

// New returns a new instance of rand.Rand which provides a random number generator using Source
// src. If the source is not available on the platform, New returns an error and *rand.Rand is nil.
// It is the responsibility of the source to be safe for concurrent use by multiple goroutines.
func New(src Source) (*rand.Rand, error) {
	// Call Assert and check if Err returns an error
	if src.Assert(); src.Err() != nil {
		// If it returns an error, then return nil and the error
		return nil, tserr.NotAvailable(&tserr.NotAvailableArgs{S: "rand source", Err: src.Err()})
	}
	// Return a new instance of rand.Rand and nil
	return rand.New(src), nil
}
