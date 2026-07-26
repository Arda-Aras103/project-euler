package main

import (
	"fmt"
)

func evenFibonacciSum(limit int) (sum int) {
	sum = 0
	prev, now := 0, 2

	for now < limit {
		sum += now
		new := 4*now + prev
		prev = now
		now = new
	}
	return
}

func main() {
	result := evenFibonacciSum(4000000)
	fmt.Println(result)
}
