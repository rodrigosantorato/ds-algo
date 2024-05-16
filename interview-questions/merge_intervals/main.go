package main

import (
	"fmt"
	"sort"
)

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

const (
	startI = iota
	endI
)

func main() {
	fmt.Println(mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) //[[1,6],[8,10],[15,18]]
	fmt.Println(mergeIntervals([][]int{{1, 4}, {4, 5}}))                    //[[1,5]]
}

func mergeIntervals(intervals [][]int) [][]int {
	var merged [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][startI] < intervals[j][startI] ||
			intervals[i][startI] == intervals[j][startI] && intervals[i][endI] < intervals[j][endI]
	})

	for i := 0; i < len(intervals); i++ {
		// no overlap
		if len(merged) == 0 || merged[len(merged)-1][endI] < intervals[i][startI] {
			merged = append(merged, intervals[i])
		}
		// overlap
		merged[len(merged)-1][endI] = max(merged[len(merged)-1][endI], intervals[i][endI])
	}

	return merged
}
