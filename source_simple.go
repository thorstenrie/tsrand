// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library packages
import (
	"math" // math
	"sync" // sync

	"github.com/thorstenrie/lpstats" // lpstats
)

// SimpleSource implements Source64 and can be used as source for a rand.Rand.
// It holds the seed s which can be initialized with Seed(int64) and a sync.Mutex
// to enable concurrent use of an instance of a SimpleSource. A SimpleSource is safe for
// concurrent use by multiple goroutines. The output might be easily predictable and is unsuitable
// for security-sensitive services.
type SimpleSource struct {
	s  int64      // seed
	mu sync.Mutex // mutex to enable concurrency
}

// NewSimpleSource returns a pointer to a new SimpleSource initiatlized with the default seed.
func NewSimpleSource() *SimpleSource {
	ex := SimpleSource{s: defaultSeed}
	return &ex
}

// Seed initializes the pseudo-random number generator source to
// a deterministic state defined by s.
func (ex *SimpleSource) Seed(s int64) {
	// Lock source
	ex.mu.Lock()
	// Set seed
	ex.s = s
	// Unlock source
	ex.mu.Unlock()
}

// Uint64 returns a pseudo-random 64-bit value. The pseudo-random value
// is calculated by two calls of Int63.
func (ex *SimpleSource) Uint64() uint64 {
	return uint64(ex.Int63())>>31 | uint64(ex.Int63())<<32
}

// Int63 returns a pseudo-random 63-bit integer determined by a
// deterministic calculation based on the seed.
func (ex *SimpleSource) Int63() int64 {
	// Lock source
	ex.mu.Lock()
	// Generate pseudo random Int63
	var (
		// Save sign of current seed
		sn float64 = float64(lpstats.Sign(ex.s))
		// Constant prime numbers p1 and p2
		p1 uint64 = 15485863   // 1000000th prime number
		p2 uint64 = 2038074743 // 100000000th prime number
		// Multiply absolute value of current seed with p1
		a uint64 = uint64(lpstats.Abs(ex.s)) * p1
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
// no used operation of SimpleSource returns an error, Err always returns nil.
func (ex *SimpleSource) Err() error {
	return nil
}

// Assert checks the availability of a random number generator source. For SimpleSource, it is empty,
// because the pseudo random number calculation is always available.
func (ex *SimpleSource) Assert() {}
