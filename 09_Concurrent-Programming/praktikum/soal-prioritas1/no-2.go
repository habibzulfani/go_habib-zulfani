package main

import (
	"fmt"
)

func main() {
	// Buat channel untuk mengirim bilangan kelipatan 3
	multiplesOfThree := make(chan int)

	// Buat goroutine untuk menghasilkan bilangan kelipatan 3 dan mengirimkannya ke channel
	go func() {
		for i := 1; ; i++ {
			if i%3 == 0 {
				multiplesOfThree <- i
			}
		}
	}()

	// Menerima dan mencetak bilangan kelipatan 3 dari channel
	for i := 0; i < 10; i++ {
		fmt.Println(<-multiplesOfThree)
	}

	// Tutup channel setelah selesai menggunakannya
	close(multiplesOfThree)
}
