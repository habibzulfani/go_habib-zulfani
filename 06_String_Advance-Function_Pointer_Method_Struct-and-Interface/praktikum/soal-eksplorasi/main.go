package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode strings.Builder

	for _, char := range s.name {
		// Encrypt uppercase letters
		if char >= 'A' && char <= 'Z' {
			encryptedChar := ((char - 'A' + 1) % 26) + 'A'
			nameEncode.WriteRune(encryptedChar)
		} else if char >= 'a' && char <= 'z' {
			// Encrypt lowercase letters
			encryptedChar := ((char - 'a' + 1) % 26) + 'a'
			nameEncode.WriteRune(encryptedChar)
		} else {
			// Keep non-alphabet characters as-is
			nameEncode.WriteRune(char)
		}
	}

	return nameEncode.String()
}

func (s *student) Decode() string {
	var nameDecode strings.Builder

	for _, char := range s.name {
		// Decrypt uppercase letters
		if char >= 'A' && char <= 'Z' {
			decryptedChar := ((char - 'A' - 1 + 26) % 26) + 'A'
			nameDecode.WriteRune(decryptedChar)
		} else if char >= 'a' && char <= 'z' {
			// Decrypt lowercase letters
			decryptedChar := ((char - 'a' - 1 + 26) % 26) + 'a'
			nameDecode.WriteRune(decryptedChar)
		} else {
			// Keep non-alphabet characters as-is
			nameDecode.WriteRune(char)
		}
	}

	return nameDecode.String()
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + " is : " + c.Decode())
	}
}