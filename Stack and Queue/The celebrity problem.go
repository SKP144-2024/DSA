package main

import (
	"fmt"
)

// Global matrix M
var M [][]int

// knows function as described in the problem
func knows(a, b int) bool {
	return M[a][b] == 1
}

// Solution struct
type Solution struct{}

// celebrity method on Solution
func (s Solution) celebrity(n int) int {
	celeb := 0
	for i := 1; i < n; i++ {
		if knows(celeb, i) {
			celeb = i
		}
	}
	for i := 0; i < n; i++ {
		if i != celeb {
			if knows(celeb, i) || !knows(i, celeb) {
				return -1
			}
		}
	}
	return celeb
}

// Helper to run test cases
func runTestCase(matrix [][]int, expected int) {
	M = matrix
	sol := Solution{}
	result := sol.celebrity(len(M))
	fmt.Println("Matrix:")
	for _, row := range M {
		fmt.Println(row)
	}
	fmt.Printf("Expected: %d, Got: %d\n", expected, result)
	if result == expected {
		fmt.Println("✅ Pass")
	} else {
		fmt.Println("❌ Fail")
	}
	fmt.Println("------------------------------------")
}

// Main function with test cases
func main() {
	runTestCase([][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 1, 0},
	}, 1)

	runTestCase([][]int{
		{0, 1, 0},
		{0, 0, 0},
		{1, 1, 0},
	}, 1)

	runTestCase([][]int{
		{0, 1, 0},
		{0, 0, 1},
		{0, 0, 0},
	}, -1)

	runTestCase([][]int{
		{0},
	}, 0)

	runTestCase([][]int{
		{0, 1},
		{0, 0},
	}, 1)

	runTestCase([][]int{
		{0, 1},
		{1, 0},
	}, -1)

	runTestCase([][]int{
		{0, 1, 1, 1},
		{0, 0, 1, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}, 3)
}
