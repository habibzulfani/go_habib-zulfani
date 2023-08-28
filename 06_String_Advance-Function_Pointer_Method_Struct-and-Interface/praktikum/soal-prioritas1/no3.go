package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var result string
	if len(a) > len(b) {
		if strings.Contains(a, b) {
			result = b
		} else {
			result = "Tidak ditemukan"
		}
	} else {
		if strings.Contains(b, a) {
			result = a
		} else {
			result = "Tidak ditemukan"
		}
	}

	return result
}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) //AKA

	fmt.Println(Compare("KANGOORO", "KANG")) //KANG

	fmt.Println(Compare("KI", "KIJANG")) //KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU

	fmt.Println(Compare("ILALANG", "ILA")) //ILA

}
