package main

import "fmt"

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func getPrime(nth int) int {
    count := 0
    num := 2
    for {
        if isPrime(num) {
            count++
        }
        if count == nth {
            return num
        }
        num++
    }
}

func main() {
    fmt.Println(isPrime(7)) // 2
    fmt.Println(getPrime(5)) // 11
    fmt.Println(getPrime(8)) // 19
    fmt.Println(getPrime(9)) // 23
    fmt.Println(getPrime(10)) // 29
}
