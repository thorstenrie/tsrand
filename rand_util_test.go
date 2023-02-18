// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library packages, lpstats and tserr
import (
	"math/rand" // rand
	"testing"   // testing

	"github.com/thorstenrie/lpstats" // lpstats
	"github.com/thorstenrie/tserr"   // tserr
)

const (
	testIntn int     = 6       // test random numbers from half-open interval [0,testIntn)
	testItr  int     = 1000000 // number of iterations for random number generation tests
	maxDiff  float64 = 0.1     // maximum difference of near equal comparison
)

// testRand tests rnd generated random numbers for integers, unsigned integers and floats.
// Each test is performed two times. For the first run, the tests are executed without setting a seed.
// For the second run, the seed is set to the default seed.
func testRand(t *testing.T, rnd *rand.Rand) {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// Test fails immediately, if rnd is nil
	if rnd == nil {
		t.Fatal(tserr.NilPtr())
	}
	// Run all tests for two iterations
	for i := 0; i <= 1; i++ {
		// Set seed to default seed for the second iteration
		if i == 1 {
			// Set seed of rnd to default seed
			rnd.Seed(defaultSeed)
		}
		// Test rnd generated random integers
		testRandInt(t, rnd)
		// Test rnd generated random floats
		testRandFloat(t, rnd)
		// Test rnd generated random unsigned integers
		testRandUint(t, rnd)
	}
}

// testRandInt retrieves random integers from rnd and stores them in a slice. The number of retrieved random numbers is defined
// by the constant testItr. The test fails, if the arithmetic mean of the random numbers or variance does not
// equal the expected values with a maximum difference defined by the constant maxDiff.
func testRandInt(t *testing.T, rnd *rand.Rand) {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// The test fails if rnd is nil
	if rnd == nil {
		t.Fatal(tserr.NilPtr())
	}
	// Allocate and initialize slice a with size testItr
	a := make([]int, testItr)
	// Iterate random number generator testItr times
	for i := 0; i < testItr; i++ {
		// Retrieve a random number from rnd in the interval [0,testIntn) and store it in the slice
		a[i] = rnd.Intn(testIntn)
	}
	// Calculate the arithmetic mean of the random integers
	mean, e := lpstats.ArithmeticMean(a)
	// The test fails if arithmetic mean has an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "a", Err: e}))
	}
	// Calculate the expected value for the interval [0, testIntn-1]
	meane := lpstats.ExpectedValueU(0, testIntn-1)
	// The test fails if the arithmetic mean does not equal the expected value with a maximum difference of maxDiff
	if !lpstats.NearEqual(mean, meane, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Arithmetic mean of a", Actual: mean, Want: meane}))
	}
	// Calculate the variance of the random integers
	vari, e := lpstats.Variance(a)
	// The test fails if variance returns an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Variance", Fn: "a", Err: e}))
	}
	// Calculate the expected variance
	varie := lpstats.VarianceN(uint(testIntn))
	// The test fails if the variance does not equal the expected variance with a maximum difference of maxDiff
	if !lpstats.NearEqual(vari, varie, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Variance of a", Actual: vari, Want: varie}))
	}
}

// testRandFloat retrieves random values as float64 of the interval [0,1) from rnd and stores them in a slice. The number of retrieved random values is defined
// by the constant testItr. The test fails, if the arithmetic mean of the random values or variance does not
// equal the expected values with a maximum difference defined by the constant maxDiff.
func testRandFloat(t *testing.T, rnd *rand.Rand) {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// The test fails if rnd is nil
	if rnd == nil {
		t.Fatal(tserr.NilPtr())
	}
	// Allocate and initialize slice a with size testItr
	a := make([]float64, testItr)
	// Iterate random number generator testItr times
	for i := 0; i < testItr; i++ {
		// Retrieve a random value from rnd in the interval [0,1) as float64 and store it in the slice
		a[i] = rnd.Float64()
	}
	// Calculate the arithmetic mean of the retrieved random values
	mean, e := lpstats.ArithmeticMean(a)
	// The test fails if arithmetic mean returns an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "a", Err: e}))
	}
	// Calculate the expected value for the interval [0,1]
	meane := lpstats.ExpectedValueU(0, 1)
	// The test fails if the arithmetic mean does not equal the expected value with a maximum difference of maxDiff
	if !lpstats.NearEqual(mean, meane, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Arithmetic mean of a", Actual: mean, Want: meane}))
	}
	// Calculate the variance of the random values
	vari, e := lpstats.Variance(a)
	// The test fails if variance returns an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Variance", Fn: "a", Err: e}))
	}
	// Calculate the expected variance
	varie := lpstats.VarianceU(0, 1)
	// The test fails if the variance does not equal the expected variance with a maximum difference of maxDiff
	if !lpstats.NearEqual(vari, varie, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Variance of a", Actual: vari, Want: varie}))
	}
}

// testRandUint retrieves random unsigned integers from rnd, normalizes them to the interval [0,1] and stores them in a slice.
// The number of retrieved random values is defined by the constant testItr. The test fails, if the arithmetic mean of the
// random values or variance does not equal the expected values with a maximum difference defined by the constant maxDiff.
func testRandUint(t *testing.T, rnd *rand.Rand) {
	// Panic if t is nil
	if t == nil {
		panic("nil pointer")
	}
	// The test fails if rnd is nil
	if rnd == nil {
		t.Fatal(tserr.NilPtr())
	}
	// Allocate and initialize slice a with size testItr
	a := make([]float64, testItr)
	// Iterate random number generator testItr times
	for i := 0; i < testItr; i++ {
		// Retrieve a random unsigned integer from rnd, normalize to [0,1] as float64 and store it in the slice
		a[i] = float64(rnd.Uint64()) / float64(^uint64(0))
	}
	// Calculate the arithmetic mean of the retrieved random values
	mean, e := lpstats.ArithmeticMean(a)
	// The test fails if the arithmetic mean has an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "a", Err: e}))
	}
	// Calculate the expected value for the interval [0,1]
	meane := lpstats.ExpectedValueU(0, 1)
	// The test fails if the arithmetic mean does not equal the expected value with a maximum difference of maxDiff
	if !lpstats.NearEqual(mean, meane, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Arithmetic mean of a", Actual: mean, Want: meane}))
	}
	// Calculate the variance of the random values
	vari, e := lpstats.Variance(a)
	// The test fails if variance returns an error
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Variance", Fn: "a", Err: e}))
	}
	// Calculate the expected variance
	varie := lpstats.VarianceU(0, 1)
	// The test fails if the variance does not equal the expected variance with a maximum difference of maxDiff
	if !lpstats.NearEqual(vari, varie, maxDiff) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "Variance of a", Actual: vari, Want: varie}))
	}
}

// benchRandUint retrieves random unsigned integers from rnd for benchmarks.
func benchRandUint(b *testing.B, rnd *rand.Rand) {
	// Panic if t is nil
	if b == nil {
		panic("nil pointer")
	}
	// The test fails if rnd is nil
	if rnd == nil {
		b.Fatal(tserr.NilPtr())
	}
	for n := 0; n < b.N; n++ {
		rnd.Uint64()
	}
}
