package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	difference := diagonalDifference(matrix)
	fmt.Println("Selisih mutlak diagonal adalah:", difference)
}

func diagonalDifference(matrix [][]int) int {
	n := len(matrix)
	leftDiagonalSum := 0
	rightDiagonalSum := 0

	for i := 0; i < n; i++ {
		leftDiagonalSum += matrix[i][i]
		rightDiagonalSum += matrix[i][n-i-1]
	}

	difference := int(math.Abs(float64(leftDiagonalSum - rightDiagonalSum)))
	return difference
}