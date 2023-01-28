// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library packages
import (
	"math" // math
	"sync" // sync
)

// simpleSource implements Source64 and can be used as source for a rand.Rand.
// It holds the seed s which can be initialized with Seed(int64) and a sync.Mutex
// to enable concurrent use of an instance of a simpleSource. A simpleSource is safe for
// concurrent use by multiple goroutines. The output might be easily predictable and is unsuitable
// for security-sensitive services.
type simpleSource struct {
	s  int64      // seed
	mu sync.Mutex // mutex to enable concurrency
}

// newSimpleSource returns a pointer to a new simpleSource initiatlized with the default seed.
func newSimpleSource() *simpleSource {
	ex := simpleSource{s: defaultSeed}
	return &ex
}

// Seed initializes the pseudo-random number generator source to
// a deterministic state defined by s.
func (ex *simpleSource) Seed(s int64) {
	// Lock source
	ex.mu.Lock()
	// Set seed
	ex.s = s
	// Unlock source
	ex.mu.Unlock()
}

// Uint64 returns a pseudo-random 64-bit value. The pseudo-random value
// is calculated by two calls of Int63.
func (ex *simpleSource) Uint64() uint64 {
	return uint64(ex.Int63())>>31 | uint64(ex.Int63())<<32
}

// Int63 returns a pseudo-random 63-bit integer determined by a
// deterministic calculation based on the seed.
func (ex *simpleSource) Int63() int64 {
	// Lock source
	ex.mu.Lock()
	// Generate pseudo random Int63
	var (
		// Save sign of current seed
		sn float64 = sign(ex.s)
		// Constant prime numbers p1 and p2
		p1 uint64 = 15485863
		p2 uint64 = 2038074743
		// Multiply absolute value of current seed with p1
		a uint64 = abs(ex.s) * p1
		// Generate a float of [0,1) and multiply with sign of current seed
		vfn float64 = sn * float64(a*a*a%p2) / float64(p2)
		// Scale to range of uint64
		vf float64 = math.Round(vfn * float64(^uint64(0)))
		// Shift to Int63
		vi int64 = int64(uint64(vf) >> 1)
	)
	// Increment seed
	ex.s++
	// Unlock source
	ex.mu.Unlock()
	// Return Int63
	return vi
}

// Err provides the last occuring error of the random number generator source. Since
// no used operation of simpleSource returns an error, Err always returns nil.
func (ex *simpleSource) Err() error {
	return nil
}

// Assert checks the availability of a random number generator source. For simpleSource it is empty,
// because the pseudo random number calculation is always available.
func (ex *simpleSource) Assert() {}
