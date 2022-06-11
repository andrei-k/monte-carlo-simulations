package main

import (
	"testing"
)

const samples = 100000000

func BenchmarkCalcPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcPi(samples)
	}
}

func BenchmarkCalcPiConcurrent(b *testing.B) {
	threads := 8
	ch := make(chan Channel, threads)
	for i := 0; i < b.N; i++ {
		calcPiConcurrent(samples, ch)
	}
}
