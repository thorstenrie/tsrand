// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tsrand

// Import standard library packages
import "math/rand" // math/rand

// Source extends the rand.Source64 interface by  Assert() and Err().
// Assert checks the availability of a random number generator source.
// A subsequent call of Err() returns an error, if the source is not
// available on the platform. Err() returns the last occuring error.
type Source interface {
	rand.Source64
	Assert()
	Err() error
}
