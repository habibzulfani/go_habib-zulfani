package main


import "fmt"


func caesar(offset int, input string) string {

// your code here
output := ""

for _, char := range input {
	if char == ' ' {
		output += " "
		continue
	}

	newASCII := (int(char)-97+offset) % 26 + 97
	output += string(rune(newASCII))
}

return output

}


func main() {

fmt.Println(caesar(3, "abc")) // def

fmt.Println(caesar(2, "alta")) // def

fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}