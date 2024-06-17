package main

import (
	"testing"
)

func Benchmark_V1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		v1()
	}
}

func Benchmark_V2(b *testing.B) {
	for i := 1; i < b.N; i++ {
		v2()
	}
}

func Benchmark_V3(b *testing.B) {
	for i := 1; i < b.N; i++ {
		v3()
	}
}
