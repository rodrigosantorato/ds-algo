package main

import "fmt"

func main() {
	array1 := []int{1, 2, 4, 7, 12}
	array2 := []int{3, 5, 8, 10, 15}

	fmt.Println(mergeTwoSortedArrays(array1, array2))
}

func mergeTwoSortedArrays(array1 []int, array2 []int) []int {
	var mergedArray []int
	firstIndex, secondIndex := 0, 0

	for firstIndex < len(array1) || secondIndex < len(array2) {
		fmt.Println(firstIndex, secondIndex)
		if firstIndex < len(array1) && (secondIndex == len(array2) || array1[firstIndex] < array2[secondIndex]) {
			mergedArray = append(mergedArray, array1[firstIndex])
			firstIndex++
			continue
		}
		mergedArray = append(mergedArray, array2[secondIndex])
		secondIndex++
	}

	return mergedArray
}
