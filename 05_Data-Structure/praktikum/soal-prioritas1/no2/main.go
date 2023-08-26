package main

import "fmt"

func Mapping(slice []string) map[string]int {
	totalStrings := make(map[string]int)
	
	for _, word := range slice {
		i := 0
		for _, existingWords := range slice{
			if word == existingWords {
				i++
			}	
		}
		
		totalStrings[word] = i
	}

	return totalStrings
}

func main() {

	// Test cases
	
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"})) // map[asd:2 qwe:1]
	
	fmt.Println(Mapping([]string{})) // map[]
	
	}