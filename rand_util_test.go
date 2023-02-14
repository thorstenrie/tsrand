package tsrand

import (
	"math/rand"
	"testing"

	"github.com/thorstenrie/lpstats"
	"github.com/thorstenrie/tserr"
)

func testRand(t *testing.T, rnd *rand.Rand) {
	if t == nil {
		panic("nil pointer")
	}
	if rnd == nil {
		t.Fatal(tserr.NilPtr())
	}
	for i := 0; i <= 1; i++ {
		if i == 1 {
			rnd.Seed(defaultSeed)
		}
		testRandInt(t, rnd)
		testRandFloat(t, rnd)
	}
}

func testRandInt(t *testing.T, rnd *rand.Rand) {
	if t == nil {
		panic("nil pointer")
	}
	if rnd == nil {
		t.Fatal(tserr.NilPtr())
	}
	a := make([]int, itr)
	for i := 0; i < itr; i++ {
		a[i] = rnd.Intn(intn)
	}
	mean, e := lpstats.ArithmeticMean(a)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "a", Err: e}))
	}
	meane := lpstats.ExpectedValueU(0, intn-1)
	if !lpstats.NearEqual(mean, meane, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Arithmetic mean of a", Actual: mean, Want: meane}))
	}
	vari, e := lpstats.Variance(a)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Variance", Fn: "a", Err: e}))
	}
	varie := lpstats.VarianceN(uint(intn))
	if !lpstats.NearEqual(vari, varie, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Variance of a", Actual: vari, Want: varie}))
	}
}

func testRandFloat(t *testing.T, rnd *rand.Rand) {
	a := make([]float64, itr)
	for i := 0; i < itr; i++ {
		a[i] = rnd.Float64()
	}
	mean, e := lpstats.ArithmeticMean(a)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "a", Err: e}))
	}
	meane := lpstats.ExpectedValueU(0, 1)
	if !lpstats.NearEqual(mean, meane, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Arithmetic mean of a", Actual: mean, Want: meane}))
	}
	vari, e := lpstats.Variance(a)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Variance", Fn: "a", Err: e}))
	}
	varie := lpstats.VarianceU(0, 1)
	if !lpstats.NearEqual(vari, varie, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Variance of a", Actual: vari, Want: varie}))
	}

}
