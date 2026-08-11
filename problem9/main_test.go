package main

import "testing"

var result int

func BenchmarkSpecialPythagorean(b *testing.B) {
	for b.Loop() {
		result = findSpecialPythagoreanTriplet(1000)
	}
}
