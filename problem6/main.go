package main

import "fmt"

func sumOfSquares(lowerLimit, upperLimit int) int {
	sum := (upperLimit*(upperLimit+1)*(2*upperLimit+1) - (lowerLimit-1)*(lowerLimit)*(2*lowerLimit-1)) / 6
	return sum
}

func squareOfSum(lowerLimit, upperLimit int) int {
	sum := ((upperLimit * (upperLimit + 1)) - (lowerLimit * (lowerLimit - 1))) / 2
	return sum * sum

}

func main() {
	fmt.Println(squareOfSum(1, 100) - sumOfSquares(1, 100))
}
