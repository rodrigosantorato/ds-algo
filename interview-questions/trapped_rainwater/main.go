package main

import (
	"fmt"
)

/*
input = [0,1,0,2,1,0,1,3,2,1,2,1]
res = 6
*/

func main() {
	fmt.Println(trappedRainWater([]int{2, 0, 1, 0, 0, 0, 1}))                      //4
	fmt.Println(trappedRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))       //6
	fmt.Println(trappedRainWater([]int{1, 0, 0, 0, 0, 0, 0, 2}))                   //6
	fmt.Println(trappedRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 1})) //7
}

func trappedRainWater(heights []int) int {
	trapped := 0
	for i, v := range heights {
		maxL, maxR := 0, 0
		j, k := i-1, i+1

		for j >= 0 {
			if heights[j] > maxL {
				maxL = heights[j]
			}
			j--
		}
		for k < len(heights) {
			if heights[k] > maxR {
				maxR = heights[k]
			}
			k++
		}

		waterI := min(maxL, maxR) - v
		if waterI > 0 {
			trapped += waterI
		}
	}
	return trapped
}
