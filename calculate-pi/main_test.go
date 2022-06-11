package main

import (
	"testing"
)

const samples = 100000000

func BenchmarkCalcPISequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcPiSequential(samples)
	}
}

func BenchmarkCalcPiConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcPiConcurrent(samples)
	}
}
