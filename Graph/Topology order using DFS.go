package main

import (
	"fmt"
)

// topoSortDFS performs Topological Sorting using DFS.
func topoSortDFS(graph map[int][]int, n int) []int {
	visited := map[int]int{}
	topo := []int{}
	hasCycle := false
	dfs := func(int) {}
	dfs = func(node int) {
		if hasCycle {
			return
		}
		if visited[node] == 2 {
			hasCycle = true
			return
		} else if visited[node] == 1 {
			return
		} else {
			visited[node] = 2
			for _, child := range graph[node] {
				if hasCycle {
					return
				}
				dfs(child)
			}
			visited[node] = 1
		}
		// fmt.Println("\t", node)
		topo = append(topo, node)
	}

	for node := range graph {
		if hasCycle {
			return []int{}
		}
		if visited[node] == 0 {
			dfs(node)
		}
	}

	fmt.Println(topo)
	for i := 0; i < len(topo)/2; i++ {
		topo[i], topo[len(topo)-1-i] = topo[len(topo)-1-i], topo[i]
	}
	return topo
}

func main() {
	// Test Case 1: Simple DAG
	graph1 := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
	}
	fmt.Println("Test 1 Output:", topoSortDFS(graph1, 6)) // Expected: Valid order e.g. [4 5 2 3 1 0]

	// Test Case 2: Linear graph
	graph2 := map[int][]int{
		0: {1},
		1: {2},
		2: {3},
	}
	fmt.Println("Test 2 Output:", topoSortDFS(graph2, 4)) // Expected: [0 1 2 3]

	// Test Case 3: Multiple valid orders
	graph3 := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
	}
	fmt.Println("Test 3 Output:", topoSortDFS(graph3, 4)) // Expected: [0 2 1 3] or [0 1 2 3]

	// Test Case 4: Graph with cycle
	graph4 := map[int][]int{
		0: {1},
		1: {2},
		2: {0},
	}
	fmt.Println("Test 4 Output:", topoSortDFS(graph4, 3)) // Expected: [] (cycle detected)

	// Test Case 5: Disconnected graph
	graph5 := map[int][]int{
		5: {},
		4: {6},
		6: {},
		3: {1},
		1: {0},
		0: {},
		2: {},
	}
	fmt.Println("Test 5 Output:", topoSortDFS(graph5, 7)) // Expected: [3 1 0 2 4 6 5] or other valid
}
