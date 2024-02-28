package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    start := time.Now()
    fmt.Println(fibonacci(43)) // Adjust the input for deeper performance analysis
    elapsed := time.Since(start)
    fmt.Printf("fibonacci took %s\n", elapsed)
}
