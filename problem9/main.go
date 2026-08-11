package main

import (
	"fmt"
)

func computeB(a, target int) int {
	if target*(target-2*a)%(2*target-2*a) != 0 {
		return 0
	}
	return target * (target - 2*a) / (2*target - 2*a)
}

func findSpecialPythagoreanTriplet(target int) int {
	for a := 1; a < target/2; a++ {
		b := computeB(a, target)
		if b == 0 {
			continue
		}

		hypo := target - a - b
		if hypo*hypo == a*a+b*b {
			return a * b * hypo
		}

	}
	return 0
}

func main() {
	fmt.Println(findSpecialPythagoreanTriplet(1000))
}
