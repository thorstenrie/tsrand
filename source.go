package tsrand

import "math/rand"

type Source interface {
	rand.Source64
	Assert()
	Err() error
}
