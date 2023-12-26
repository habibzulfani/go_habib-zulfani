package main

import (
	"fmt"
	"time"
)

func main() {
	// Buat channel dengan buffer untuk mengirim bilangan kelipatan 3
	multiplesOfThree := make(chan int, 10) // Buffer channel dengan kapasitas 10

	// Buat goroutine untuk menghasilkan bilangan kelipatan 3 dan mengirimkannya ke channel
	go func() {
		for i := 1; ; i++ {
			if i%3 == 0 {
				multiplesOfThree <- i
			}
		}
	}()

	// Menerima dan mencetak bilangan kelipatan 3 dari channel
	for {
		select {
		case multiple := <-multiplesOfThree:
			fmt.Println(multiple)
		case <-time.After(1 * time.Second):
			fmt.Println("Waktu habis. Menghentikan program.")
			return
		}
	}
}
