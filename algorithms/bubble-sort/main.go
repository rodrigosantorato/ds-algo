package main

import "fmt"

func main() {
	arr := []int{9, 4, 6, 99, 42, 965, 2, 7, 4, 1, 14, 17, 54, 87}

	fmt.Println(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j+1 != length && arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
