// fibonacci.go

package main

import "fmt"

// fibonacci returns a slice containing the Fibonacci series up to n terms.
func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	series := make([]int, n)
	series[0] = 0
	if n > 1 {
		series[1] = 1
		for i := 2; i < n; i++ {
			series[i] = series[i-1] + series[i-2]
		}
	}
	return series
}

func main() {
	n := 10 // Change this value to generate more or fewer terms
	fmt.Printf("Fibonacci series up to %d terms: %v\n", n, fibonacci(n))
}