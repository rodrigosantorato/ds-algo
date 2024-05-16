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

	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	var merged []int
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) || rightIndex < len(right) {
		if leftIndex < len(left) && rightIndex == len(right) || leftIndex < len(left) && left[leftIndex] < right[rightIndex] {
			merged = append(merged, left[leftIndex])
			leftIndex++
		} else {
			merged = append(merged, right[rightIndex])
			rightIndex++
		}
	}
	return merged
}
