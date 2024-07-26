
# Go Programs Collection

#### This repository contains various Go programs that I've learned and implemented. Each program demonstrates different aspects of Go programming, from basic syntax to more advanced concepts.

## Factorial Program
```go
package main

import "fmt"

var n = 5

func main() {
	fmt.Printf("Factorial of %d is %d ", n, factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

```
#### Output
```
Factorial of 5 is 120 
Program exited.
```