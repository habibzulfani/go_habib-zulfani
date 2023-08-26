package main

import "fmt"
	
func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}

	i := 5
	for i*i <= number {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))          // true
	fmt.Println(primeNumber(17))          // true
	fmt.Println(primeNumber(20))          // false
	fmt.Println(primeNumber(35))          // false
}
