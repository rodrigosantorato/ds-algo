package main

import "fmt"

func main() {
	fmt.Println(findFactorialRecursive(4))
	fmt.Println(findFactorialIterative(4))
}

func findFactorialRecursive(n int) int {
	if n > 0 {
		return n * findFactorialRecursive(n-1)
	}
	return 1
}

func findFactorialIterative(n int) int {
	factorial := n
	for i := n - 1; i > 0; i-- {
		factorial = factorial * i
	}
	return factorial
}
