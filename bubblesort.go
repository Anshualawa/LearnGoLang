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
