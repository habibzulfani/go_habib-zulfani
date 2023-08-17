package main

import "fmt"

func main() {
	n := 25
	fmt.Printf("angka : %d \n", n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		} 
	}
}
