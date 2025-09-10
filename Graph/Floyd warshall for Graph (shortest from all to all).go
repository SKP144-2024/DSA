package main

import (
	"fmt"
	"math"
	"reflect"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func floydWarshall(n int, edges [][3]int) [][]int {
	// TODO: implement Floyd-Warshall
	// return 2D matrix of shortest paths between all pairs
	pathMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		pathMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				pathMatrix[i][j] = math.MaxInt32
			}
		}
	}

	for _, edge := range edges {
		pathMatrix[edge[0]][edge[1]] = edge[2]
	}

	// fmt.Println(pathMatrix)

	isNegCycled := false
	for via := 0; via < n; via++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// fmt.Println(i, via, j, " => ", pathMatrix[i][j], pathMatrix[i][via], pathMatrix[via][j], pathMatrix[i][via]+pathMatrix[via][j])
				if pathMatrix[i][via] != math.MaxInt32 && pathMatrix[via][j] != math.MaxInt32 {
					pathMatrix[i][j] = min(pathMatrix[i][j], pathMatrix[i][via]+pathMatrix[via][j])
				}
				if i == j && pathMatrix[i][j] < 0 {
					isNegCycled = true
				}
			}
		}
		// fmt.Println(pathMatrix)
	}

	if isNegCycled {
		return [][]int{}
	}

	return pathMatrix
}

func main() {
	tests := []struct {
		n      int
		edges  [][3]int
		output [][]int
	}{
		// Test 1: Graph with a negative cycle (0→3→0 cycle = -2)
		{
			4,
			[][3]int{
				{0, 1, 3}, {0, 2, 8}, {0, 3, -4},
				{1, 3, 7}, {1, 2, 1},
				{2, 1, 4},
				{3, 0, 2}, {3, 2, -5},
			},
			[][]int{}, // negative cycle → return empty
		},

		// Test 2: Graph with negative edge (no negative cycle)
		{
			3,
			[][3]int{
				{0, 1, 4}, {0, 2, 11},
				{1, 2, -2},
			},
			[][]int{
				{0, 4, 2},
				{math.MaxInt32, 0, -2},
				{math.MaxInt32, math.MaxInt32, 0},
			},
		},

		// Test 3: Disconnected graph
		{
			3,
			[][3]int{
				{0, 1, 5},
			},
			[][]int{
				{0, 5, math.MaxInt32},
				{math.MaxInt32, 0, math.MaxInt32},
				{math.MaxInt32, math.MaxInt32, 0},
			},
		},

		// Test 4: Single node graph
		{
			1,
			[][3]int{},
			[][]int{
				{0},
			},
		},
	}

	for i, test := range tests {
		result := floydWarshall(test.n, test.edges)
		pass := reflect.DeepEqual(result, test.output)
		if pass {
			fmt.Printf("✅ Test %d Passed: got %v, expected %v\n", i+1, result, test.output)
		} else {
			fmt.Printf("❌ Test %d Failed: got %v, expected %v\n", i+1, result, test.output)
		}
	}
}
