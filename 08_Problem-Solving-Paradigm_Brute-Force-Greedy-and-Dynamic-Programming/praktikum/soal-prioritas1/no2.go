package main

import "fmt"

func generate(numRows int) [][]int {
	// Inisialisasi matriks segitiga Pascal
	pascalTriangle := make([][]int, numRows)
	for i := range pascalTriangle {
		pascalTriangle[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		// Set nilai pertama dan terakhir selalu 1
		pascalTriangle[i][0] = 1
		pascalTriangle[i][i] = 1

		// Hitung nilai-nilai di antara nilai pertama dan terakhir
		for j := 1; j < i; j++ {
			pascalTriangle[i][j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}
	}

	return pascalTriangle
}

func main() {
	numRows := 5
	result := generate(numRows)
	fmt.Println(result)
}
