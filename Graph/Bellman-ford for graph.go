package main

import (
	"fmt"
	"math"
)

type Edge struct {
	u, v, w int
}

func bellmanFord(n int, edges []Edge, src int) ([]int, bool) {
	// TODO: implement Bellman-Ford
	// return (distances, hasNegativeCycle)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0
	for i := 1; i < n; i++ {
		for _, edge := range edges {
			src, dest, weight := edge.u, edge.v, edge.w
			if dist[dest] > dist[src]+weight {
				dist[dest] = dist[src] + weight
			}
		}
	}
	// fmt.Println(dist)
	isNegCycled := false
	for _, edge := range edges {
		src, dest, weight := edge.u, edge.v, edge.w
		if dist[dest] > dist[src]+weight {
			isNegCycled = true
			break
		}
	}

	if isNegCycled {
		return []int{}, isNegCycled
	} else {
		return dist, isNegCycled
	}
}

func main() {
	tests := []struct {
		n        int
		edges    []Edge
		src      int
		expected []int
		hasCycle bool
	}{
		// Test 1: Simple positive edges
		{
			5,
			[]Edge{
				{0, 1, 6}, {0, 2, 7},
				{1, 2, 8}, {1, 3, 5}, {1, 4, -4},
				{2, 3, -3}, {2, 4, 9},
				{3, 1, -2},
				{4, 0, 2}, {4, 3, 7},
			},
			0,
			[]int{0, 2, 7, 4, -2},
			false,
		},
		// Test 2: Graph with negative edge but no cycle
		{
			4,
			[]Edge{
				{0, 1, 4},
				{0, 2, 5},
				{1, 2, -3},
				{2, 3, 2},
			},
			0,
			[]int{0, 4, 1, 3},
			false,
		},
		// Test 3: Graph with reachable negative cycle
		{
			3,
			[]Edge{
				{0, 1, 1},
				{1, 2, -1},
				{2, 0, -1},
			},
			0,
			[]int{}, // no valid distances
			true,
		},
		// Test 4: Disconnected graph
		{
			4,
			[]Edge{
				{0, 1, 3},
			},
			0,
			[]int{0, 3, math.MaxInt32, math.MaxInt32},
			false,
		},
		// Test 5: The graph from your image
		{
			5,
			[]Edge{
				{0, 1, 2},
				{1, 2, -2},
				{2, 0, -1},
				{1, 3, 3},
				{3, 4, -8},
				{4, 2, 1},
			},
			0,
			[]int{}, // negative cycle
			true,
		},
	}

	for i, test := range tests {
		dist, hasCycle := bellmanFord(test.n, test.edges, test.src)
		if hasCycle != test.hasCycle || (!hasCycle && !equal(dist, test.expected)) {
			fmt.Printf("Test %d: FAIL ❌ (got %v, %v expected %v, %v)\n", i+1, dist, hasCycle, test.expected, test.hasCycle)
		} else {
			fmt.Printf("Test %d: PASS ✅ \n", i+1)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
