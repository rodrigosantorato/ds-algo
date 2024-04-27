package main

import "fmt"

/**
Find container with most water and return the area of said container
each number in the input array is a height, and the width is the diff from one index to the other

scenario 1:
biggest container is 9 to 7, the height of the container is 7 and the width is 5
i.e. (index of 7 = 5) - (index of 9 = 0) = 5
input: [9,1,2,3,4,7] res: 35
**/

func main() {
	fmt.Println(findAreaOfBiggestContainer([]int{9, 1, 2, 3, 4, 7}))
	fmt.Println(findAreaOfBiggestContainer([]int{9, 30, 2, 3, 50, 7}))
}

func findAreaOfBiggestContainer(nums []int) int {
	// compare pairs of numbers and check their area as containers
	// get note of the max area and return after all the comparisons
	maxArea := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// do the comparisons
			// calc the area of the container
			// grab the min height and multiply by the diff in index
			area := min(nums[i], nums[j]) * (j - i)
			if area > maxArea {
				maxArea = area
			}
		} // checked every single comparison for index i
	}

	return maxArea
}
