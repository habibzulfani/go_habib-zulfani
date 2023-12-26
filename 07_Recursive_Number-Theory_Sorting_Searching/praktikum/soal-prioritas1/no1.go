package main

import "fmt"

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}

/* 5 = 5
	fibo(4) + fibo(3) 
		fibo(3)+fibo(2)	 +	fibo(2) + fibo(1) 
			fibo(2)+fibo(1) + fibo(1)+fibo(0) + fibo(1)+fibo(0) + 1
			fibo(1)+fibo(0) + 1 + 1 + 1 + 1 + 1 + 1
			1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 = 1 	
*/
func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(5))  // 8
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
