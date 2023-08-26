package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	counts := make(map[int]int) // Map untuk menghitung kemunculan setiap angka
	seen := make(map[int]bool)   // Map untuk melacak angka yang sudah terlihat

	var hasil []int

	for _, r := range angka {
		num, _ := strconv.Atoi(string(r))
		if seen[num] {
			counts[num] = 0 // Jika angka sudah terlihat sebelumnya, set count menjadi 0
		} else {
			counts[num]++
			seen[num] = true
		}
	}

	for _, r := range angka {
		num, _ := strconv.Atoi(string(r))
		if counts[num] == 1 {
			hasil = append(hasil, num)
		}
	}

	return hasil

	// var numberSlice []int
	// for _, numStr := range angka {
	// 	num, err := strconv.Atoi(string(numStr))
	// 	if err != nil {
	// 		fmt.Printf("Error: %v\n", err)
	// 		continue
	// 	}
	// 	numberSlice = append(numberSlice, num)
	// }
	
	// totalStrings := make(map[int]int)
	
	// for _, a := range numberSlice {
	// 	i := 0
	// 	for _, b := range numberSlice{
	// 		if a == b {
	// 			i++
	// 		}	
	// 	}
		
	// 	totalStrings[a] = i
	// }

	// var hasil []int
	// for num, count := range totalStrings {
	// 	if count == 1 {
	// 		hasil = append(hasil, num)
	// 	}
	// }

	// return hasil
}

func main() {

	// Test cases
	
	fmt.Println(munculSekali("1234123")) // [4]
	
	fmt.Println(munculSekali("76523752")) // [6 3]
	
	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]
	
	fmt.Println(munculSekali("1122334455")) // []
	
	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]
	
	}