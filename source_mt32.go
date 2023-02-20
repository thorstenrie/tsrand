// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// MT32Source implements Source64 and can be used as source for a rand.Rand. It is based on the
// reference implementation of the 32-bit Mersenne Twister.
// MT32Source holds the pseudo-random number generator internal states mt and mti. A MT32Source is not safe for
// concurrent use by multiple goroutines. The output might be easily predictable and is unsuitable
// for security-sensitive services.
type MT32Source struct {
	mt  []uint32 // slice for the state vector
	mti int      // index and mti==N+1 means mt is not initialized
}

// Period parameters based on the reference implementation of the 32-bit Mersenne Twister
var (
	mt32c = struct {
		n, m                               int
		matrixA, defaultSeed, uMask, lMask uint32
	}{
		n:           624,
		m:           397,
		matrixA:     0x9908b0df, // constant vector a
		defaultSeed: 5489,       // default seed
		uMask:       0x80000000, // most significant w-r bits
		lMask:       0x7fffffff, // least significant r bits
	}
)

// NewMT32Source returns a new instance of MT32Source. MT32Source implements Source64,
// is based on the reference implementation of the 32-bit Mersenne Twister and can be used as source for a rand.Rand.
// A MT32Source is not safe for concurrent use by multiple goroutines. The output might be
// easily predictable and is unsuitable for security-sensitive services.
func NewMT32Source() *MT32Source {
	src := &MT32Source{mt: make([]uint32, mt32c.n), mti: mt32c.n + 1}
	return src
}

// seed initializes the state vector with seed s.
func (src *MT32Source) seed(s int64) {
	src.mt[0] = uint32(s & 0xffffffff)
	for src.mti = 1; src.mti < mt32c.n; src.mti++ {
		src.mt[src.mti] = (uint32(1812433253)*(src.mt[src.mti-1]^(src.mt[src.mti-1]>>30)) + uint32(src.mti))
	}
}

// Seed initializes the state vector with seed s.
func (src *MT32Source) Seed(s int64) {
	// Initialization of the state vector with seed s
	src.seed(s)
}

// uint32 returns a pseudo-random 32-bit value. The implementation is
// based on mt19937.ar.c
func (src *MT32Source) uint32() uint32 {
	var (
		y     uint32
		mag01 [2]uint32 = [2]uint32{0, mt32c.matrixA}
	)
	if src.mti >= mt32c.n {
		var kk int
		// Initialize state vector with default seed if not initialized with a seed before
		if src.mti == mt32c.n+1 {
			src.seed(int64(mt32c.defaultSeed))
		}
		for kk = 0; kk < mt32c.n-mt32c.m; kk++ {
			y = (src.mt[kk] & mt32c.uMask) | (src.mt[kk+1] & mt32c.lMask)
			src.mt[kk] = src.mt[kk+mt32c.m] ^ (y >> 1) ^ mag01[y&0x1]
		}
		for ; kk < mt32c.n-1; kk++ {
			y = (src.mt[kk] & mt32c.uMask) | (src.mt[kk+1] & mt32c.lMask)
			src.mt[kk] = src.mt[kk+(mt32c.m-mt32c.n)] ^ (y >> 1) ^ mag01[y&0x1]
		}
		y = (src.mt[mt32c.n-1] & mt32c.uMask) | (src.mt[0] & mt32c.lMask)
		src.mt[mt32c.n-1] = src.mt[mt32c.m-1] ^ (y >> 1) ^ mag01[y&0x1]
		src.mti = 0
	}
	y = src.mt[src.mti]
	src.mti += 1
	// Tempering
	y ^= (y >> 11)
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= (y >> 18)
	return y
}

// Uint64 returns a pseudo-random 64-bit value. The pseudo-random value
// is calculated by two calls of uint32().
func (src *MT32Source) Uint64() uint64 {
	return uint64(src.uint32()) | uint64(src.uint32())<<32
}

// Int63 returns a pseudo-random 63-bit integer. The pseudo-random value
// is calculated by two calls of uint32().
func (src *MT32Source) Int63() int64 {
	return int64(src.Uint64() >> 1)
}

// Err provides the last occurring error of the random number generator source. Since
// no used operation of MT32Source returns an error, Err always returns nil.
func (src *MT32Source) Err() error {
	return nil
}

// Assert checks the availability of a random number generator source. For MT32Source, it is empty,
// because the pseudo random number calculation is always available.
func (src *MT32Source) Assert() {}
