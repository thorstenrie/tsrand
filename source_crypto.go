package tsrand

import (
	crand "crypto/rand"
	"encoding/binary"
	"sync"
)

type cSource struct {
	mu sync.Mutex
	e  error
}

var (
	cryptoSrc cSource
)

func cryptoSource() *cSource {
	return &cryptoSrc
}

func (c *cSource) Seed(s int64) {}

func (c *cSource) Uint64() (v uint64) {
	c.mu.Lock()
	c.e = binary.Read(crand.Reader, binary.BigEndian, &v)
	c.mu.Unlock()
	return v
}

func (c *cSource) Int63() int64 {
	vu := c.Uint64()
	mask := ^uint64(1 << 63)
	return int64(vu & mask)
}

func (c *cSource) Assert() {
	b := make([]byte, 1)
	c.mu.Lock()
	_, c.e = crand.Read(b)
	c.mu.Unlock()
}

func (c *cSource) Err() error {
	c.mu.Lock()
	e := c.e
	c.mu.Unlock()
	return e

}
