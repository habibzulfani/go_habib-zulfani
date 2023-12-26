package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var sharedValue int

	// Jumlah goroutine yang akan dijalankan
	numGoroutines := 5

	// Wait Group untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Fungsi yang akan dijalankan oleh setiap goroutine
	goroutineFunc := func(id int) {
		defer wg.Done()

		// Mengunci mutex sebelum mengakses dan memodifikasi sharedValue
		mutex.Lock()
		sharedValue++
		fmt.Printf("Goroutine %d: Nilai bersama setelah penambahan: %d\n", id, sharedValue)
		mutex.Unlock() // Membuka mutex setelah selesai

		// Di sini Anda dapat melakukan operasi lain yang tidak memerlukan mutex
	}

	// Menjalankan beberapa goroutine
	for i := 0; i < numGoroutines; i++ {
		go goroutineFunc(i)
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	fmt.Printf("Nilai bersama akhir: %d\n", sharedValue)
}
