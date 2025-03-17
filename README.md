
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
## Sorting Program
```go
package main

import "fmt"

func main() {
	arr := []int{3, 4, 5, 2, 1,9,8,6,5,4,1,3,16,5,4}

	fmt.Println("Original array:", arr)
	fmt.Println("Sorted array:", Sort(arr))
}

func Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

```
#### Output
```
Original array: [3 4 5 2 1]
Sorted array: [1 2 3 4 5]

Program exited.
```

## Revers sorted Array
```go
package main
import "fmt"

func reverse(arr []int) []int {
	left,right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func main(){
	array := []int{4,5,6,7,8,9}
	fmt.Println("Original array :",array)
	fmt.Println("Reverse array :",reverse(array))
}
```
### Output
```
Original array : [4,5,6,7,8,9]
Reverse array :  [9,8,7,6,5,4]
```
