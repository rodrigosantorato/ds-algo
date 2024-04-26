package main

import "fmt"

func main() {
	fmt.Println(findTwoSumBruteForce([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(findTwoSumBruteForce([]int{1, 2, 10, 3, 4, 5, 20}, 30))
	fmt.Println(findTwoSumHashMap([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(findTwoSumHashMap([]int{1, 2, 10, 3, 4, 5, 20}, 30))
}

func findTwoSumBruteForce(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	i, j := 0, 1
	for i < len(nums) {
		// if we're done checking with no answer we increment i and go back to the beginning with the j checks
		if j >= len(nums) {
			i++
			j = i + 1
			continue
		}
		// check each number if summed with nums[i] will result in the target
		if nums[i]+nums[j] == target {
			return []int{i, j}
		}
		j++
	}

	return []int{}
}

func findTwoSumHashMap(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	// key = number to be found | val = index of the number that satisfies "number + number to be found = target"
	numbersToFind := make(map[int]int)

	for i, v := range nums {
		if numbersToFind[v] != 0 {
			return []int{numbersToFind[v], i}
		}
		numbersToFind[target-v] = i
	}
	return []int{}
}
