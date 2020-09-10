package main

import (
	"testing"
)

func BenchmarkFrequencyByMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencyByMap()
	}
}

func BenchmarkFrequencyByArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencyByArray()
	}
}

func BenchmarkCheapReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cheapReplace()
	}
}

func BenchmarkExpensiveReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expensiveReplace()
	}
}
