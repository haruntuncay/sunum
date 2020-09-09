package main

import (
	"testing"
)

func BenchmarkDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo()
	}
}

func BenchmarkDemo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo2()
	}
}
