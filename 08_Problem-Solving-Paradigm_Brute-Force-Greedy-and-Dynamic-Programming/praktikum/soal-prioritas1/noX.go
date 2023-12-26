package main

import (
	"fmt"
	// "math"
)

func main() {
	n := 24
	// m := int(math.Sqrt(float64(n)))
	m := n/2
	for i := 1; i<= m; i++ {
		if n % i == 0 {
			fmt.Println(i)
		}
	}
	
}