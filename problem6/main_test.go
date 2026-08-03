package main

import "testing"

var result int

func BenchmarkDifference(b *testing.B) {
	for b.Loop() {
		result = squareOfSum(1, 100) - sumOfSquares(1, 100)
	}
}
