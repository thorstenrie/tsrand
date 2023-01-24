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
	rnd, err := CryptoRand()
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
	rnd, _ := New(NewExampleSource())
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(1)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(2)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(3)
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
	rnd.Seed(time.Now().UnixNano())
	fmt.Println([3]int{rnd.Intn(6) + 1, rnd.Intn(6) + 1, rnd.Intn(6) + 1})
}
