// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

import "sync"

type MT64Source struct {
	mt  []uint64
	mti int
	mu  sync.Mutex // mutex to enable concurrency
}

var (
	mt64c = struct {
		n, m                               int
		matrixA, defaultSeed, uMask, lMask uint64
	}{
		n:           312,
		m:           156,
		matrixA:     0xB5026F5AA96619E9,
		defaultSeed: 5489,
		uMask:       0xFFFFFFFF80000000,
		lMask:       0x7FFFFFFF,
	}
)

func NewMT64Source() *MT64Source {
	src := &MT64Source{mt: make([]uint64, mt64c.n), mti: mt64c.n + 1}
	return src
}

// Unsafe for concurrency
func (src *MT64Source) seedUnsafe(s int64) {
	src.mt[0] = uint64(s)
	for src.mti = 1; src.mti < mt64c.n; src.mti++ {
		src.mt[src.mti] = (uint64(6364136223846793005)*(src.mt[src.mti-1]^(src.mt[src.mti-1]>>62)) + uint64(src.mti))
	}
}

func (src *MT64Source) Seed(s int64) {
	// Lock source
	src.mu.Lock()
	src.seedUnsafe(s)
	// Unlock source
	src.mu.Unlock()
}

func (src *MT64Source) Uint64() uint64 {
	var (
		x     uint64
		mag01 [2]uint64 = [2]uint64{0, mt64c.matrixA}
	)
	// Lock source
	src.mu.Lock()
	if src.mti >= mt64c.n {
		var i int
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

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	return x
}

func (src *MT64Source) Int63() int64 {
	return int64(src.Uint64() >> 1)
}

func (src *MT64Source) Err() error {
	return nil
}

func (src *MT64Source) Assert() {}
