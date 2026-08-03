package main

import (
	"fmt"
)

func divisibleByK(coefficient int, limit int) (sum int) {
	var numberOfTerm = (limit - 1) / coefficient
	sum = coefficient * numberOfTerm * (numberOfTerm + 1) / 2
	return
}

func main() {
	result := divisibleByK(3, 1000) + divisibleByK(5, 1000) - divisibleByK(15, 1000)
	fmt.Println(result)
}
