package main

import "fmt"

func sieveOfEratosthenes(limit int) []int {
	primeList := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		primeList[i] = true
	}
	factor := 2

	for factor*factor <= limit {
		if primeList[factor] {
			for current := factor * factor; current <= limit; current += factor {
				primeList[current] = false
			}
		}
		factor++
	}

	result := []int{}
	for current := 2; current <= limit; current++ {
		if primeList[current] {
			result = append(result, current)
		}
	}

	return result
}

func sumOfPrimes(list []int) (sum int) {
	for _, number := range list {
		sum += number
	}
	return
}

func main() {
	result := sieveOfEratosthenes(2000000)
	fmt.Println(sumOfPrimes(result))
}
