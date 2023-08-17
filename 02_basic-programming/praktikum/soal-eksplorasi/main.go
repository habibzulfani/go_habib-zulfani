package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Masukkan sebuah kata: ")
	fmt.Scan(&input)

	if isPalindrome(input) {
		fmt.Printf("%s adalah palindrome.\n", input)
	} else {
		fmt.Printf("%s bukan palindrome.\n", input)
	}
}

func isPalindrome(str string) bool {
	str = strings.ToLower(str) // Mengubah huruf menjadi lowercase
	str = strings.ReplaceAll(str, " ", "") // Menghilangkan spasi dari kata

	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}