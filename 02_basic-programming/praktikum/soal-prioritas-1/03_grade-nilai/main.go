package main

import "fmt"

func main() {
	nilai := 57
	var grade string

	if nilai >= 0 && nilai <= 100 {
		if nilai >= 80 {
			grade = "A"
		} else if nilai >= 65 {
			grade = "B"
		} else if nilai >= 50 {
			grade = "C"
		} else if nilai >= 35 {
			grade = "D"
		} else {
			grade = "E"
		}

		fmt.Printf("Grade anda adalah: %s\n", grade)
	} else {
		fmt.Println("Nilai Invalid")
	}
}