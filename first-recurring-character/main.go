package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 1, 8, 3, 5, 8}

	fmt.Println(findFirstRecurringElement(arr))
}

func findFirstRecurringElement(arr []int) int {
	seenCharacters := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		if seenCharacters[arr[i]] {
			return arr[i]
		}

		seenCharacters[arr[i]] = true
	}

	return 0
}
