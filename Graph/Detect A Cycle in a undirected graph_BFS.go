package main

import (
	"fmt"
)

type node struct {
	val    int
	parent int
}

// TODO: Implement this function
func hasCycleDFS(graph map[int][]int) bool {
	visited := map[int]bool{}
	for nod := range graph {
		if visited[nod] {
			continue
		}
		queue := []node{node{val: nod, parent: -1}}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			visited[curr.val] = true
			for _, child := range graph[curr.val] {
				if !visited[child] {
					queue = append(queue, node{child, curr.val})

				} else {
					if child != curr.parent {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	testCases := []struct {
		graph     map[int][]int
		expect    bool
		visualize string
	}{
		{
			graph: map[int][]int{
				0: {1},
				1: {0},
			},
			expect:    false,
			visualize: "Simple Edge: 0 -- 1",
		},
		{
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
			},
			expect:    true,
			visualize: "Triangle Cycle: 0 -- 1 -- 2 -- 0",
		},
		{
			graph: map[int][]int{
				0: {1},
				1: {0, 2},
				2: {1, 3},
				3: {2},
			},
			expect:    false,
			visualize: "Straight Line: 0 -- 1 -- 2 -- 3",
		},
		{
			graph: map[int][]int{
				0: {1},
				1: {0, 2, 3},
				2: {1, 3},
				3: {1, 2},
			},
			expect:    true,
			visualize: "Square with diagonal: 1-2-3 forms a cycle",
		},
		{
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 3},
				2: {0, 3, 4},
				3: {1, 2, 5},
				4: {2, 6},
				5: {3, 6},
				6: {4, 5},
			},
			expect:    true,
			visualize: "Complex Multi-Cycle: 2-4-6-5-3-2 and others",
		},
		{
			graph: map[int][]int{
				0: {1},
				1: {0, 2},
				2: {1},
				3: {4},
				4: {3},
				5: {},
			},
			expect:    false,
			visualize: "Disconnected Components: No cycle anywhere",
		},
		{
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 3, 4},
				2: {0, 5},
				3: {1},
				4: {1, 5},
				5: {2, 4},
				6: {7},
				7: {6},
			},
			expect:    true,
			visualize: "Disconnected + One Component Cycle: 4-5-2-0-1-4",
		},
		{
			graph: map[int][]int{
				0: {1},
				1: {0, 2},
				2: {1, 3},
				3: {2},
				4: {5, 6},
				5: {4, 6},
				6: {4, 5},
			},
			expect:    true,
			visualize: "Disconnected: Component {4,5,6} has cycle (4-5-6-4)",
		},
	}

	for i, tc := range testCases {
		got := hasCycleDFS(tc.graph)
		status := "❌"
		if got == tc.expect {
			status = "✅"
		}
		fmt.Printf("Test Case %d: %v\n", i+1, tc.visualize)
		fmt.Printf("Expected: %v, Got: %v %s\n\n", tc.expect, got, status)
	}
}
