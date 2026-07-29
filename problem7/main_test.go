package main

import "testing"

var result int

func BenchmarkFindPrime(b *testing.B) {
	for b.Loop() {
		findPrime(10001)
	}
}
