package main

import "fmt"

func main() {
	arr := [16]int{9, 10, 11, 4, 6, 99, 42, 965, 2, 7, 4, 1, 14, 17, 54, 87}

	fmt.Println(insertionSort(arr))
}

func insertionSort(arr [16]int) [16]int {
	length := len(arr)
	placeholder := 0
	for i := 1; i < length; i++ {
		if arr[i] < arr[0] {
			placeholder = arr[i]
			copy(arr[1:i+1], arr[0:i])
			arr[0] = placeholder
		}
		for j := 1; j < i; j++ {
			if arr[j-1] <= arr[i] && arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
