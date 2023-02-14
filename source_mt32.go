// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

import "sync"

type MT32Source struct {
	mt  []uint32
	mti int
	mu  sync.Mutex // mutex to enable concurrency
}

var (
	mt32c = struct {
		n, m                               int
		matrixA, defaultSeed, uMask, lMask uint32
	}{
		n:           624,
		m:           397,
		matrixA:     0x9908b0df,
		defaultSeed: 5489,
		uMask:       0x80000000,
		lMask:       0x7fffffff,
	}
)

func NewMT32Source() *MT32Source {
	src := &MT32Source{mt: make([]uint32, mt32c.n), mti: mt32c.n + 1}
	return src
}

func (src *MT32Source) seedUnsafe(s int64) {
	src.mt[0] = uint32(s & 0xffffffff)
	for src.mti = 1; src.mti < mt32c.n; src.mti++ {
		src.mt[src.mti] = (uint32(1812433253)*(src.mt[src.mti-1]^(src.mt[src.mti-1]>>30)) + uint32(src.mti))
	}
}

func (src *MT32Source) Seed(s int64) {
	// Lock source
	src.mu.Lock()
	src.seedUnsafe(s)
	// Unlock source
	src.mu.Unlock()
}

func (src *MT32Source) uint32() uint32 {
	var (
		y     uint32
		mag01 [2]uint32 = [2]uint32{0, mt32c.matrixA}
	)
	// Lock source
	src.mu.Lock()
	if src.mti >= mt32c.n {
		var kk int
		if src.mti == mt32c.n+1 {
			src.seedUnsafe(int64(mt32c.defaultSeed))
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

	// Unlock source
	src.mu.Unlock()

	y ^= (y >> 11)
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= (y >> 18)

	return y
}

func (src *MT32Source) Uint64() uint64 {
	return uint64(src.uint32()) | uint64(src.uint32())<<32
}

func (src *MT32Source) Int63() int64 {
	return int64(src.Uint64() >> 1)
}

func (src *MT32Source) Err() error {
	return nil
}

func (src *MT32Source) Assert() {}
