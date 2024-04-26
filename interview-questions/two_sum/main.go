package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(twoSum([]int{1, 2, 10, 3, 4, 5, 20}, 30))
}

func twoSum(arr []int, target int) []int {
	if len(arr) < 2 {
		return []int{}
	}

	i, j := 0, 1
	for i < len(arr) {
		// if we're done checking with no answer we increment i and go back to the beginning with the j checks
		if j >= len(arr) {
			i++
			j = i + 1
			continue
		}
		// check each number if summed with arr[i] will result in the target
		if arr[i]+arr[j] == target {
			return []int{i, j}
		}
		j++
	}

	return []int{}
}
