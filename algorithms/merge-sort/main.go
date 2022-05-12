package main

import "fmt"

func main() {
	arr := []int{9, 10, 11, 4, 6, 99, 42, 965, 2, 7, 4, 1, 14, 17, 54, 87}

	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2

	left := arr[:mid]
	right := arr[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var leftIndex, rightIndex int
	var mergedArray []int

	for leftIndex < len(left) || rightIndex < len(right) {
		if leftIndex < len(left) && (rightIndex == len(right) || left[leftIndex] < right[rightIndex]) {
			mergedArray = append(mergedArray, left[leftIndex])
			leftIndex++
			continue
		}
		mergedArray = append(mergedArray, right[rightIndex])
		rightIndex++
		continue
	}
	return mergedArray
}
