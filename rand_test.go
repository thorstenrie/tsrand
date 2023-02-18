// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library package testing and tserr
import (
	"testing" // testing

	"github.com/thorstenrie/tserr" // tserr
)

// Each Test function generates (pseudo-)random numbers using the defined source of the test. It generates random values of types integer, unsigned integer, and float64.
// The Test functions compare for each type the arithmetic mean and variance of the retrieved random numbers with the expected values for mean and variance.
// If the arithmetic mean and variance of the retrieved random numbers differ more than the constant maxDiff from expected values, the test fails.
// Therefore, the Test functions provide an indication if the sources for random number generators are providing random values in expected boundaries.
// The Test functions do not evaluate the quality of retrieved random numbers and implementation of the random number generator source. The output of the random number
// generator sources might be easily predictable and unsuitable for security-sensitive services.

// TestCryptoRand retrieves random values from the cryptographically secure random number generator and performs the defined tests on arithmetic mean and variance.
// The test fails, if the cryptographically secure random number generator is not available on the platform or if tests on the retrieved random numbers fail.
func TestCryptoRand(t *testing.T) {
	// Retrieve the cryptographically secure random number generator
	rnd, err := NewCryptoRand()
	// The test fails if an error occurs
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "CryptoRand", Err: err}))
	}
	// Perform tests on the random number generator source
	testRandInt(t, rnd)
	testRandFloat(t, rnd)
	testRandUint(t, rnd)
}

// BenchmarkCryptoRand performs a benchmark on the cryptographically secure random number generator
func BenchmarkCryptoRand(b *testing.B) {
	// Retrieve the cryptographically secure random number generator
	rnd, err := NewCryptoRand()
	// The test fails if an error occurs
	if err != nil {
		b.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "CryptoRand", Err: err}))
	}
	benchRandUint(b, rnd)
}

// TestPseudoRand retrieves random values from the pseudo-random number generator and performs the defined tests on arithmetic mean and variance.
// The test fails, if the pseudo-random number generator is not available on the platform or if tests on the retrieved random numbers fail.
func TestPseudoRand(t *testing.T) {
	// Retrieve the pseudo-random number generator
	rnd, err := NewPseudoRandomRand()
	// The test fails if an error occurs
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewPseudoRandomRand", Err: err}))
	}
	// Perform tests on the random number generator source
	testRand(t, rnd)
}

// BenchmarkPseudoRand performs a benchmark on the pseudo-random number generator
func BenchmarkPseudoRand(b *testing.B) {
	// Retrieve the pseudo-random number generator
	rnd, err := NewPseudoRandomRand()
	// The test fails if an error occurs
	if err != nil {
		b.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewPseudoRandomRand", Err: err}))
	}
	benchRandUint(b, rnd)
}

// TestDeterministicRand retrieves random values from the deterministic pseudo-random number generator and performs the defined tests on arithmetic mean and variance.
// The test fails, if the deterministic pseudo-random number generator is not available on the platform or if tests on the retrieved random numbers fail.
func TestDeterministicRand(t *testing.T) {
	// Retrieve the deterministic pseudo-random number generator
	rnd, err := NewDeterministicRand()
	// The test fails if an error occurs
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewDeterministicRand", Err: err}))
	}
	// Perform tests on the random number generator source
	testRand(t, rnd)
}

// BenchmarkDeterministicRand performs a benchmark on the deterministic pseudo-random number generator
func BenchmarkDeterministicRand(b *testing.B) {
	// Retrieve the pseudo-random number generator
	rnd, err := NewDeterministicRand()
	// The test fails if an error occurs
	if err != nil {
		b.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewDeterministicRand", Err: err}))
	}
	benchRandUint(b, rnd)
}

// TestSimpleRand retrieves random values from a very simple implementation of an example pseudo-random number generator source
// and performs the defined tests on arithmetic mean and variance. The test fails, if the simple pseudo-random number generator
// is not available on the platform  or if tests on the retrieved random numbers fail.
func TestSimpleRand(t *testing.T) {
	// Retrieve the very simple pseudo-random number generator
	rnd, err := New(NewSimpleSource())
	// The test fails if an error occurs
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewSimpleSource", Err: err}))
	}
	// Perform tests on the random number generator source
	testRand(t, rnd)
}

// BenchmarkSimpleRand performs a benchmark on the very simple implementation of an example pseudo-random number generator source
func BenchmarkSimpleRand(b *testing.B) {
	// Retrieve the pseudo-random number generator
	rnd, err := New(NewSimpleSource())
	// The test fails if an error occurs
	if err != nil {
		b.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewSimpleSource", Err: err}))
	}
	benchRandUint(b, rnd)
}

// TestMT32Rand retrieves random values from an implementation based on the 32-bit Mersenne Twister
// and performs the defined tests on arithmetic mean and variance. The test fails, if the pseudo-random number generator
// is not available on the platform  or if tests on the retrieved random numbers fail.
func TestMT32Rand(t *testing.T) {
	// Retrieve the pseudo-random number generator
	rnd, err := New(NewMT32Source())
	// The test fails if an error occurs
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewMT32Source", Err: err}))
	}
	// Perform tests on the random number generator source
	testRand(t, rnd)
}

// BenchmarkMT32Rand performs a benchmark on the on the 32-bit Mersenne Twister based implemented pseudo-random number generator
func BenchmarkMT32Rand(b *testing.B) {
	// Retrieve the pseudo-random number generator
	rnd, err := New(NewMT32Source())
	// The test fails if an error occurs
	if err != nil {
		b.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewMT32Source", Err: err}))
	}
	benchRandUint(b, rnd)
}

// TestMT64Rand retrieves random values from an implementation based on the 64-bit Mersenne Twister
// and performs the defined tests on arithmetic mean and variance. The test fails, if the pseudo-random number generator
// is not available on the platform  or if tests on the retrieved random numbers fail.
func TestMT64Rand(t *testing.T) {
	// Retrieve the pseudo-random number generator
	rnd, err := New(NewMT64Source())
	// The test fails if an error occurs
	if err != nil {
		t.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewMT64Source", Err: err}))
	}
	// Perform tests on the random number generator source
	testRand(t, rnd)
}

// BenchmarkMT64Rand performs a benchmark on the on the 64-bit Mersenne Twister based implemented pseudo-random number generator
func BenchmarkMT64Rand(b *testing.B) {
	// Retrieve the pseudo-random number generator
	rnd, err := New(NewMT64Source())
	// The test fails if an error occurs
	if err != nil {
		b.Fatal(tserr.NotAvailable(&tserr.NotAvailableArgs{S: "NewMT64Source", Err: err}))
	}
	benchRandUint(b, rnd)
}
