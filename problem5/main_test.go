package main

import "testing"

func BenchmarkFindSmallestDivisibleNumber(b *testing.B) {
	for b.Loop() {
		findSmallestDivisibleNumber(1, 20)
	}
}
