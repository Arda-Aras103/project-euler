package main

import (
	"fmt"
)

func isPalindrome(number int) bool {
	reverse := 0
	temp := number

	for temp != 0 {
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}

	return reverse == number
}

func findLargestPalindrome(lower, upper int) (result int) {
	largestPalindrome := 0
	for i := upper; i >= lower; i-- {
		for j := i; j >= lower; j-- {
			product := i * j
			if isPalindrome(product) && largestPalindrome < product {
				largestPalindrome = product
			}
		}
	}
	return largestPalindrome
}

func main() {
	fmt.Println(findLargestPalindrome(100, 999))
}
