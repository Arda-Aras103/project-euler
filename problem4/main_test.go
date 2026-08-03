package main

import "testing"

func BenchmarkLargestPalindrome(b *testing.B) {
	for b.Loop() {
		findLargestPalindrome(100, 999)
	}
}
