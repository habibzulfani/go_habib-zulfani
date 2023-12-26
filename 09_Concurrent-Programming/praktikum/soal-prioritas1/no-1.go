package main

import (
    "fmt"
    "time"
)

func printMultiples(x, interval int) {
    for i := 1; ; i++ {
        if i%x == 0 {
            fmt.Printf("%d adalah kelipatan dari %d\n", i, x)
        }
        time.Sleep(time.Duration(interval) * time.Second)
    }
}

func main() {
    x := 9 
    interval := 3 

    go printMultiples(x, interval)

    // Biarkan program berjalan selama beberapa saat agar goroutine dapat bekerja
    time.Sleep(27 * time.Second)
}
