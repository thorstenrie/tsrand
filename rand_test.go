// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

import (
	"fmt"
	"testing"
	"time"

	"github.com/thorstenrie/tserr"
)

func TestMRand(t *testing.T) {
	rnd, _ := NewDeterministicRand()
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(1)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(2)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(time.Now().UnixNano())
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(time.Now().UnixNano())
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
}

func TestCrypto(t *testing.T) {
	rnd, err := NewCryptoRand()
	if err != nil {
		t.Error(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "Crypto Rand", Err: err}))
	}
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(1)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
}

func TestExRand(t *testing.T) {
	rnd, _ := New(newSimpleSource())
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(10)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(11)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(12)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(time.Now().UnixNano())
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
}
