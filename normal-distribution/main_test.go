package main

import (
	"testing"
)

func BenchmarkRunConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runConcurrent()
	}
}

func BenchmarkRunSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runSequential()
	}
}
