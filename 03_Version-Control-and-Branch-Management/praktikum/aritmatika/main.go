package main

import (
	"fmt"
	"math"
)

func main() {
    // Input nilai
    var angka1, angka2 float64

    fmt.Print("Masukkan angka pertama: ")
    fmt.Scanln(&angka1)

    fmt.Print("Masukkan angka kedua: ")
    fmt.Scanln(&angka2)

    // Operasi aritmatika
    tambah := angka1 + angka2
    kurang := angka1 - angka2
    kali := angka1 * angka2
    bagi := angka1 / angka2
	
	
	

    // Menampilkan hasil
    fmt.Printf("Hasil Penjumlahan: %.2f\n", tambah)
    fmt.Printf("Hasil Pengurangan: %.2f\n", kurang)
    fmt.Printf("Hasil Perkalian: %.2f\n", kali)
    fmt.Printf("Hasil Pembagian: %.2f\n", bagi)
	
	//konversi float to int
	i := math.Round(angka1)
	j := math.Round(angka2)

	k  := int(i)
	l  := int(j)

	// Memeriksa angka pertama
	if k%2 == 0 {
		fmt.Printf("%d adalah angka genap.\n", k)
	} else {
		fmt.Printf("%d adalah angka ganjil.\n", k)
	}
	
	// Memeriksa angka kedua
	if l%2 == 0 {
		fmt.Printf("%d adalah angka genap.\n", l)
	} else {
		fmt.Printf("%d adalah angka ganjil.\n", l)
	}
	

	
	
	
	
	
	
}
