// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// sign returns -1 as float64, if a is negative, +1 as float64
// if positive.
func sign(a int64) float64 {
	// Bitmask for signbit
	var mask uint64 = uint64(1 << 63)
	// Check if signbit is set to 1
	if uint64(a)&mask == mask {
		// Return -1 as float64
		return float64(-1)
	}
	// Return +1 as float64 otherwise
	return float64(1)

}

// abs returns the absolute value fo a as uint64.
// Note: The smallest value of a int64 does not have a
// matching positive value. The abs function returns a
// negative value in this case.
func abs(a int64) uint64 {
	// Return -a as uint64 if a is negative
	if a < 0 {
		return uint64(-a)
	}
	// Return +a as uint64 otherwise
	return uint64(a)
}
