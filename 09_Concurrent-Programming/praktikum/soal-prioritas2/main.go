package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	// Membagi teks menjadi beberapa bagian (sesuai dengan jumlah goroutine)
	numGoroutines := 4
	textParts := make([]string, numGoroutines)

	partSize := len(text) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := (i + 1) * partSize
		if i == numGoroutines-1 {
			end = len(text)
		}
		textParts[i] = text[start:end]
	}

	// Menggunakan Wait Group untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Membuat map untuk menghitung frekuensi huruf
	frequencyMap := make(map[rune]int)

	// Fungsi yang akan dijalankan oleh setiap goroutine
	goroutineFunc := func(part string) {
		defer wg.Done()
		for _, char := range part {
			// Hanya memproses huruf alfabet
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
				frequencyMap[char]++
			}
		}
	}

	// Menjalankan goroutine untuk setiap bagian teks
	for i := 0; i < numGoroutines; i++ {
		go goroutineFunc(textParts[i])
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	// Mencetak frekuensi huruf
	for char, count := range frequencyMap {
		fmt.Printf("%c : %d\n", char, count)
	}
}
