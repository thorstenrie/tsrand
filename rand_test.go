// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

import (
	"testing"

	"github.com/thorstenrie/tserr"
)

const (
	intn    int     = 20      // test random numbers from half-open interval [0,intn)
	itr     int     = 1000000 // number of iterations for random number generation tests
	maxDiff float64 = 0.1
)

func TestCryptoRand(t *testing.T) {
	rnd, err := NewCryptoRand()
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "CryptoRand", Err: err}))
	}
	if rnd == nil {
		t.Fatal(tserr.NilPtr())
	}
	testRandInt(t, rnd)
	testRandFloat(t, rnd)
}

func TestPseudoRand(t *testing.T) {
	rnd, err := NewPseudoRandomRand()
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewPseudoRandomRand", Err: err}))
	}
	testRand(t, rnd)
}

func TestDeterministicRand(t *testing.T) {
	rnd, err := NewDeterministicRand()
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewDeterministicRand", Err: err}))
	}
	testRand(t, rnd)
}

func TestSimpleRand(t *testing.T) {
	rnd, err := New(NewSimpleSource())
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewSimpleSource", Err: err}))
	}
	testRand(t, rnd)
}

func TestMT32Rand(t *testing.T) {
	rnd, err := New(NewMT32Source())
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewMT32Source", Err: err}))
	}
	testRand(t, rnd)
}

func TestMT64Rand(t *testing.T) {
	rnd, err := New(NewMT64Source())
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewMT64Source", Err: err}))
	}
	testRand(t, rnd)
}
