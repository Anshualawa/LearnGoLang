package main

import "fmt"

var n = 5  // enter number

func main() {
	fmt.Printf("Factorial of %d is %d \n", n, factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

