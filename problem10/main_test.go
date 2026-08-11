package main

import (
	"testing"
)

var result int

func Benchmark(b *testing.B) {
	for b.Loop() {
		list := sieveOfEratosthenes(2000000)
		result = sumOfPrimes(list)
	}
}
