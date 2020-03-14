package main

import (
	"math"
	"testing"
)

//  go test -bench . *.go

func BenchmarkSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Sin(.5)
	}
}
