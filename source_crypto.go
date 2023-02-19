// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library packages
import (
	crand "crypto/rand" // crypto/rand
	"encoding/binary"   // encoding/binary
	"sync"              // sync
)

// cSource implements Source64 and can be used as source for a rand.Rand.
// cSource calls crypto/rand and therefore, it provides a cryptographically secure
// random number generator source. To check, if it is available on the platform
// Assert() should be called. If it is available, Err() will return nil, otherwise
// will return an error. It holds error e which contains the last occurring error, if any,
// and a sync.Mutex to enable concurrent use. cSource is safe for concurrent use by multiple goroutines.
// cSource cannot be seeded, therefore Seed(int64) is empty.
type cSource struct {
	mu sync.Mutex // mutex to enable concurrency
	e  error      // last error occurring, if any
}

// cryptoSrc is the cSource instance used by the interface
var (
	cryptoSrc cSource
)

// cryptoSource returns the used cSource as pointer
func cryptoSource() *cSource {
	return &cryptoSrc
}

// cSource cannot be seeded. Seed(int64) is empty for cSource.
func (c *cSource) Seed(s int64) {}

// Uint64 returns a random 64-bit value.
func (c *cSource) Uint64() (v uint64) {
	// Lock source
	c.mu.Lock()
	// Read from crypto/rand provided Reader in v
	c.e = binary.Read(crand.Reader, binary.BigEndian, &v)
	// Unlock source
	c.mu.Unlock()
	// Return v as uint64
	return v
}

// Int63 returns a random 63-bit integer
func (c *cSource) Int63() int64 {
	// Retrieve a random 64-bit value with Uint64() in vu
	vu := c.Uint64()
	// Bitmask for the first 63 bits
	mask := ^uint64(1 << 63)
	// Return the first 63 bits of vu
	return int64(vu & mask)
}

// Assert checks the availability of a random number generator source.
// A subsequent call of Err() returns an error, if the source is not
// available on the platform.
func (c *cSource) Assert() {
	// Create []byte of size 1
	b := make([]byte, 1)
	// Lock the source
	c.mu.Lock()
	// Read from crypto/rand in b.
	_, c.e = crand.Read(b)
	// Unlock the source
	c.mu.Unlock()
}

// Err provides the last occurring error of the random number generator source, if any.
// It returns nil, if no error occurrred.
func (c *cSource) Err() error {
	// Lock the source
	c.mu.Lock()
	// Set return value e to struct e
	e := c.e
	// Unlock the source
	c.mu.Unlock()
	// Return e
	return e
}
