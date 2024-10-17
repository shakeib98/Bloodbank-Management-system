package main

import "fmt"

// Function to generate Fibonacci sequence up to n terms
func fibonacci(n int) []int {
    fibSeq := make([]int, n)
    fibSeq[0], fibSeq[1] = 0, 1
    for i := 2; i < n; i++ {
        fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
    }
    return fibSeq
}

func main() {
    var n int
    fmt.Print("Enter the number of terms: ")
    fmt.Scan(&n)

    if n < 2 {
        fmt.Println("Please enter a number greater than or equal to 2")
        return
    }

    fibSeq := fibonacci(n)
    fmt.Println("Fibonacci sequence:", fibSeq)
}