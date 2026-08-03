package main

import "fmt"

func gcd(first, second int) int {
	for first != 0 {
		second, first = first, second%first
	}
	return second
}

func lcm(first, second int) int {
	return (first * second) / gcd(first, second)
}

func findSmallestDivisibleNumber(lowerLimit, upperLimit int) (result int) {
	result = lowerLimit
	for i := lowerLimit + 1; i <= upperLimit; i++ {
		result = lcm(result, i)
	}
	return
}

func main() {
	fmt.Println(findSmallestDivisibleNumber(1, 20))
}
