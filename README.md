
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
## Find Maximum and Minimum in an Array
```go
package main
import "fmt"

func MaxMinValue(array []int)(int,int){
	if len(array) == 0 {
		fmt.Println("Array is empty!")
		return 0,0
	}
	min,max := array[0], array[0]
	for _,num := range array {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main () {
	array := []int{9,7,8,4,3,2,1,5}
	fmt.Println("Original Array :",array)

	min, max := MaxMinValue(array)

	fmt.Println("Minimum value :",min)
	fmt.Println("Maximum value :",max)
}
```
## Merge Two Sorted Arrays
```go
package main
import "fmt"

func Mrg2SrtArray(array1, array2 []int) []int {
	i,j := 0,0
	merged := []int{}

	for i < len(array1) && j < len(array2) {
		if array1[i] < array2[j] {
			merged = append(merged,array1[i])
			i++
		} else {
			merged = append(merged, array2[j])
			j++
		}
	}
	for i < len(array1) {
		merged = append(merged, array1[i])
		i++
	}
	for j < len(array2) {
		merged = append(merged,array2[j])
		j++
	}
	return merged
}

func main(){
	arr1 := []int{3,4,5,6}
	arr2 := []int{8,9,10,13,16}
	fmt.Println("Array1 :",arr1)
	fmt.Println("Array2 :",arr2)
	fmt.Println("Merged array :",Mrg2SrtArray(arr1,arr2))
}
```
