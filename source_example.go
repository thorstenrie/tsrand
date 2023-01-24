package tsrand

import (
	"math"
	"sync"
)

type exSource struct {
	s  int64
	mu sync.Mutex
}

const (
	defaultexSeed int64 = 1
)

func NewExampleSource() *exSource {
	ex := exSource{s: defaultexSeed}
	return &ex
}

func (ex *exSource) Seed(s int64) {
	ex.mu.Lock()
	ex.s = s
	ex.mu.Unlock()
}

func (ex *exSource) Uint64() uint64 {
	return uint64(ex.Int63())>>31 | uint64(ex.Int63())<<32
}

func (ex *exSource) Int63() int64 {
	ex.mu.Lock()
	a := ex.s * 15485863
	vf := float64(a*a*a%2038074743) / 2038074743 // [0,1)]
	vi := int64(uint64(math.Round(vf*float64(^uint64(0)))) >> 1)
	ex.s++
	ex.mu.Unlock()
	return vi
}

func (ex *exSource) Err() error {
	return nil
}

func (ex *exSource) Assert() {}
