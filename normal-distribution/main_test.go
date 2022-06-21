package main

import (
	"testing"
)

func BenchmarkRunSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runSequential()
	}
}

func BenchmarkRunConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runConcurrent()
	}
}
