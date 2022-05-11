package main

import "fmt"

func main() {
	arr := []int{9, 4, 6, 99, 42, 965, 2, 7, 4, 1, 14, 17, 54, 87}

	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	length := len(arr)
	minIndex := 0
	for i := 0; i < length; i++ {
		minIndex = i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
