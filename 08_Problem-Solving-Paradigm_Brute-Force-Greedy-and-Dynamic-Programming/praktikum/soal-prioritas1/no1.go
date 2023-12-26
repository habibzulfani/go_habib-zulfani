package main

import "fmt"

func intToBinary(n int) int {
	if n == 0 {
		return 0
	}

	binaryValue := 0
	bitPosition := 1

	for n > 0 {
		remainder := n % 2
		binaryValue += remainder * bitPosition
		n = n/2
		bitPosition *= 10
	}

	return binaryValue
}

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = intToBinary(i)
	}
	return ans
}

func main() {
	n := 3
	result := countBits(n)
	fmt.Println(result)
	fmt.Println(countBits(n))
}
