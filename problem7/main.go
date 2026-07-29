package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number <= 3 {
		return number > 1
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	for current := 5; current*current <= number; current += 6 {
		if number%current == 0 || number%(current+2) == 0 {
			return false
		}
	}

	return true
}
func findPrime(number int) int {
	primeCounter := 1
	finalPrime := 2
	for current := 3; primeCounter != number; current += 2 {
		if isPrime(current) {
			finalPrime = current
			primeCounter++
		}
	}
	return finalPrime
}

func main() {
	fmt.Println(findPrime(10001))
}
