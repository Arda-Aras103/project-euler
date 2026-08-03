package main

import (
	"fmt"
)

func findLargestPrimeFactor(number int64) (maxPrime int64) {
	maxPrime = -1

	for number%2 == 0 {
		maxPrime = 2
		number /= 2
	}

	for number%3 == 0 {
		maxPrime = 3
		number /= 3
	}

	var index int64 = 5
	for index*index <= number {
		for number%index == 0 {
			maxPrime = index
			number /= index
		}
		for number%(index+2) == 0 {
			maxPrime = index + 2
			number /= (index + 2)
		}
		index += 6
	}

	if number > 1 {
		maxPrime = number
	}

	return

}

func main() {
	result := findLargestPrimeFactor(600851475143)
	fmt.Println(result)
}
