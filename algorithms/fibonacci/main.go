package main

import "fmt"

func main() {
	fmt.Println(findFibonacciIterative(9))
	fmt.Println(findFibonacciSecondSolution(9))
	fmt.Println(findFibonacciRecursive(9))
}

func findFibonacciIterative(n int) int {
	placeholder, previousFibonacci, currentFibonacci := 0, 0, 1

	for i := 1; true; i++ {
		if i == n {
			return currentFibonacci
		}
		placeholder = previousFibonacci
		previousFibonacci = currentFibonacci
		currentFibonacci = currentFibonacci + placeholder
	}
	return 0
}

func findFibonacciSecondSolution(n int) int {
	slice := []int{0, 1}

	for i := 2; i <= n; i++ {
		slice = append(slice, slice[i-1]+slice[i-2])
	}
	return slice[n]
}

func findFibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return findFibonacciRecursive(n-1) + findFibonacciRecursive(n-2)
}
