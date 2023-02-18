# tsrand
Go package for random numbers with a simple API

[![Go Report Card](https://goreportcard.com/badge/github.com/thorstenrie/tsrand)](https://goreportcard.com/report/github.com/thorstenrie/tsrand)
[![CodeFactor](https://www.codefactor.io/repository/github/thorstenrie/tsrand/badge)](https://www.codefactor.io/repository/github/thorstenrie/tsrand)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorstenrie/tsrand)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorstenrie/tsrand)](https://pkg.go.dev/mod/github.com/thorstenrie/tsrand)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorstenrie/tsrand)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorstenrie/tsrand)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorstenrie/tsrand)
![GitHub last commit](https://img.shields.io/github/last-commit/thorstenrie/tsrand)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorstenrie/tsrand)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorstenrie/tsrand)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorstenrie/tsrand)
![GitHub](https://img.shields.io/github/license/thorstenrie/tsrand)

The package tsrand provides a simple interface for random numbers. Each interface function returns a [rnd.Rand](https://pkg.go.dev/math/rand#Rand) for a specified random number generator. A returned rnd.Rand instance uses the specified random number generator to provide random numbers over its interface. The package exposes the random number generators [math/rand](https://pkg.go.dev/math/rand) and [crypto/rand](https://pkg.go.dev/crypto/rand) from the Go standard library as well as builtin example random number generators [SimpleSource](https://pkg.go.dev/github.com/thorstenrie/tsrand#SimpleSource), [MT32Source](https://pkg.go.dev/github.com/thorstenrie/tsrand#MT32Source), and [MT64Source](https://pkg.go.dev/github.com/thorstenrie/tsrand#MT64Source). Also, the interface enables the use of a custom random number generator source with function [New](https://pkg.go.dev/github.com/thorstenrie/tsrand#New).

- **Simple**: Without configuration, just function calls
- **Easy to use**: Retrieve random numbers with [rnd.Rand](https://pkg.go.dev/math/rand#Rand)
- **Tested**: Unit tests with high [code coverage](https://gocover.io/github.com/thorstenrie/tsrand)
- **Dependencies**: Only depends on the [Go Standard Library](https://pkg.go.dev/std), [tserr](https://github.com/thorstenrie/tserr) and [lpstats](https://github.com/thorstenrie/lpstats)

## Usage

The package is installed with 

```
go get github.com/thorstenrie/tsrand
```

In the Go app, the package is imported with

```
import "github.com/thorstenrie/tsrand"
```
## Random number generators

Each interface function returns a [rnd.Rand](https://pkg.go.dev/math/rand#Rand) for a specified random number generator. A returned rnd.Rand instance uses the specified random number generator to provide random numbers over its interface.

- Cryptographically secure random number generator based on [crypto/rand](https://pkg.go.dev/crypto/rand)
- Pseudo-random number generator based on [math/rand](https://pkg.go.dev/math/rand)
- Deterministic pseudo-random number generator based on [math/rand](https://pkg.go.dev/math/rand)
- A custom implementation of a random number generator [Source](https://pkg.go.dev/github.com/thorstenrie/tsrand#Source) with [New](https://pkg.go.dev/github.com/thorstenrie/tsrand#New)
- Example of a very simple pseudo-random number generator [SimpleSource](https://pkg.go.dev/github.com/thorstenrie/tsrand#SimpleSource) based on an very simple example from [Wikipedia](https://en.wikipedia.org/wiki/Pseudorandom_number_generator#Implementation)
- Example pseudo-random number generator [MT32Source](https://pkg.go.dev/github.com/thorstenrie/tsrand#MT32Source) based on the [32-bit Mersenne Twister](http://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/MT2002/emt19937ar.html)
- Example pseudo-random number generator [MT64Source](https://pkg.go.dev/github.com/thorstenrie/tsrand#MT64Source) based on the [64-bit Mersenne Twister](http://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/emt64.html)

Except for the cryptographically secure random number generator, the output of the pseudo-random number generators might be easily predictable and is unsuitable for security-sensitive services.

## Unit tests

Each Test function generates (pseudo-)random numbers using the defined source of the test. It generates random values of types integer, unsigned integer, and float64. The Test functions compare for each type the arithmetic mean and variance of the retrieved random numbers with the expected values for mean and variance. If the arithmetic mean and variance of the retrieved random numbers differ more than the constant maxDiff from expected values, the test fails. Therefore, the Test functions provide an indication if the sources for random number generators are providing random values in expected boundaries. The Test functions do not evaluate the quality of retrieved random numbers and implementation of the random number generator source. The output of the random number generator sources might be easily predictable and unsuitable for security-sensitive services.

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/tsrand)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/tsrand)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Ftsrand)
