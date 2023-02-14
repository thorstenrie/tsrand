// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library packages
import (
	mrand "math/rand" // math/rand
	"sync"            // sync
)

// dSource implements Source64, uses the math/rand pseudo-random number generator and can be used as source for a rand.Rand.
// It holds the deterministic pseudo-random number generator drnd and a sync.Mutex
// to enable concurrent use of an instance of a dSource. A dSource is safe for
// concurrent use by multiple goroutines. The output might be easily predictable and is unsuitable
// for security-sensitive services.
type dSource struct {
	mu   sync.Mutex  // mutex to enable concurrency
	drnd *mrand.Rand // deterministic pseudo-random number generator
}

// NewDeterministicSource returns a new instance of dSource. Source implements Source64,
// uses the math/rand pseudo-random number generator and can be used as source for a rand.Rand.
// A dSource is safe for concurrent use by multiple goroutines. The output might be
// easily predictable and is unsuitable for security-sensitive services. The deterministic
// pseud-random number generator is initiatlized with seed.
func newDeterministicSource(seed int64) *dSource {
	// Create new instance of dSource in ms
	ms := dSource{}
	// Create a new rand.Rand in drnd
	ms.drnd = mrand.New(mrand.NewSource(seed))
	// Return ms
	return &ms
}

// Seed initializes the pseudo-random number generator source to
// a deterministic state defined by s.
func (m *dSource) Seed(s int64) {
	// Lock the source
	m.mu.Lock()
	// Set seed pf drnd to s
	m.drnd.Seed(s)
	// Unlock the source
	m.mu.Unlock()
}

// Uint64 returns a pseudo-random 64-bit value.
func (m *dSource) Uint64() uint64 {
	// Lock the source
	m.mu.Lock()
	// Retrieve pseudo-random uint64 from drnd in v
	v := m.drnd.Uint64()
	// Unlock the source
	m.mu.Unlock()
	// Return v
	return v
}

// Int63 returns a random 63-bit integer
func (m *dSource) Int63() int64 {
	// Lock the source
	m.mu.Lock()
	// Retrieve pseudo-random 64-bit integer from drnd in v
	v := m.drnd.Int63()
	// Unlock the source
	m.mu.Unlock()
	// Return v
	return v
}

// Err provides the last occuring error of the random number generator source. Since
// no used operation of dSource returns an error, Err always returns nil.
func (m *dSource) Err() error {
	return nil
}

// Assert checks the availability of a random number generator source. For dSource it is empty,
// because the pseudo random number calculation is always available.
func (m *dSource) Assert() {}
