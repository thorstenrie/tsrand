package tsrand

import (
	"fmt"
	"math/rand"
)

type realnum interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func testRandInt(rnd *rand.Rand) {
	var a [itr]int
	for i := 0; i < itr; i++ {
		a[i] = rnd.Intn(intn)
	}
	mean := mean(a)
	meane := float64(intn-1) / float64(2)
	vari := variance(a, mean)
	varie := (square(float64(intn)) - 1) / float64(12)
	fmt.Printf("I) Actual mean = %f, Expected mean = %f\n", mean, meane)
	fmt.Printf("I) Actual variance = %f, Expected variance = %f\n", vari, varie)
}

func testRandFloat(rnd *rand.Rand) {
	var a [itr]float64
	for i := 0; i < itr; i++ {
		a[i] = rnd.Float64()
	}
	mean := mean(a)
	meane := 0.5
	vari := variance(a, mean)
	varie := float64(1) / float64(12)
	fmt.Printf("F) Actual mean = %f, Expected mean = %f\n", mean, meane)
	fmt.Printf("F) Actual variance = %f, Expected variance = %f\n", vari, varie)
}

func mean[T realnum](x [itr]T) float64 {
	var sum float64
	for i := 0; i < itr; i++ {
		sum += float64(x[i])
	}
	return sum / float64(itr)
}

func variance[T realnum](x [itr]T, mean float64) float64 {
	var vari float64
	for i := 0; i < itr; i++ {
		vari += square(float64(x[i])-mean) / float64(itr)
	}
	return vari
}

func square(x float64) float64 {
	return x * x
}
