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
	array := []int{1,2,3,4,5,6}

	fmt.Println("original array :",array)
	fmt.Println("reversed array :",reverse(array))
}