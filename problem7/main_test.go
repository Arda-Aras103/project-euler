package main

import "testing"

var result int

func BenchmarkFindPrime(b *testing.B) {
	for b.Loop() {
		result = findPrime(10001)
	}
}
