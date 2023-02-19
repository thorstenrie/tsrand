// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library package sync
import "sync" // sync

// MT64Source implements Source64 and can be used as source for a rand.Rand. It is based on the
// reference implementation of the 64-bit Mersenne Twister.
// MT64Source holds the pseudo-random number generator internal states mt and mti and a sync.Mutex
// to enable concurrent use of an instance of a MT64Source. A MT64Source is safe for
// concurrent use by multiple goroutines. The output might be easily predictable and is unsuitable
// for security-sensitive services.
type MT64Source struct {
	mt  []uint64   // slice for the state vector
	mti int        // index and mti==N+1 means mt is not initialized
	mu  sync.Mutex // mutex to enable concurrency
}

// Period parameters based on the reference implementation of the 64-bit Mersenne Twister
var (
	mt64c = struct {
		n, m                               int
		matrixA, defaultSeed, uMask, lMask uint64
	}{
		n:           312,
		m:           156,
		matrixA:     0xB5026F5AA96619E9, // constant vector a
		defaultSeed: 5489,               // default seed
		uMask:       0xFFFFFFFF80000000, // most significant 33 bits
		lMask:       0x7FFFFFFF,         // least significant 31 bits
	}
)

// NewMT64Source returns a new instance of MT64Source. MT64Source implements Source64,
// is based on the reference implementation of the 64-bit Mersenne Twister and can be used as source for a rand.Rand.
// A MT64Source is safe for concurrent use by multiple goroutines. The output might be
// easily predictable and is unsuitable for security-sensitive services.
func NewMT64Source() *MT64Source {
	src := &MT64Source{mt: make([]uint64, mt64c.n), mti: mt64c.n + 1}
	return src
}

// seedUnsafe initializes the state vector with seed s. It is not safe for concurrent use.
func (src *MT64Source) seedUnsafe(s int64) {
	src.mt[0] = uint64(s)
	for src.mti = 1; src.mti < mt64c.n; src.mti++ {
		src.mt[src.mti] = (uint64(6364136223846793005)*(src.mt[src.mti-1]^(src.mt[src.mti-1]>>62)) + uint64(src.mti))
	}
}

// Seed initializes the state vector with seed s. It is safe for concurrent use.
func (src *MT64Source) Seed(s int64) {
	// Lock source
	src.mu.Lock()
	// Initialization of the state vector with seed s
	src.seedUnsafe(s)
	// Unlock source
	src.mu.Unlock()
}

// Uint64 returns a pseudo-random 64-bit value. The implementation is
// based on mt19937-64.c
func (src *MT64Source) Uint64() uint64 {
	var (
		x     uint64
		mag01 [2]uint64 = [2]uint64{0, mt64c.matrixA}
	)
	// Lock source
	src.mu.Lock()
	if src.mti >= mt64c.n {
		var i int
		// Initialize state vector with default seed if not initialized with a seed before
		if src.mti == mt64c.n+1 {
			src.seedUnsafe(int64(mt64c.defaultSeed))
		}
		for i = 0; i < mt64c.n-mt64c.m; i++ {
			x = (src.mt[i] & mt64c.uMask) | (src.mt[i+1] & mt64c.lMask)
			src.mt[i] = src.mt[i+mt64c.m] ^ (x >> 1) ^ mag01[x&1]
		}
		for ; i < mt64c.n-1; i++ {
			x = (src.mt[i] & mt64c.uMask) | (src.mt[i+1] & mt64c.lMask)
			src.mt[i] = src.mt[i+(mt64c.m-mt64c.n)] ^ (x >> 1) ^ mag01[x&1]
		}
		x = (src.mt[mt64c.n-1] & mt64c.uMask) | (src.mt[0] & mt64c.lMask)
		src.mt[mt64c.n-1] = src.mt[mt64c.m-1] ^ (x >> 1) ^ mag01[x&1]
		src.mti = 0
	}
	x = src.mt[src.mti]
	src.mti += 1
	// Unlock source
	src.mu.Unlock()
	// Tempering
	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	return x
}

// Int63 returns a pseudo-random 64-bit integer.
func (src *MT64Source) Int63() int64 {
	return int64(src.Uint64() >> 1)
}

// Err provides the last occurring error of the random number generator source. Since
// no used operation of MT64Source returns an error, Err always returns nil.
func (src *MT64Source) Err() error {
	return nil
}

// Assert checks the availability of a random number generator source. For MT64Source, it is empty,
// because the pseudo random number calculation is always available.
func (src *MT64Source) Assert() {}
