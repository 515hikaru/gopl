package main

import "fmt"

func reverse(arr *[10]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", arr)
	ptr := &arr
	reverse(ptr)
	fmt.Printf("%v\n", arr)
}
