package main

import "fmt"

// topoSortBFS should return a topological order of the graph if possible.
// If a cycle exists, return an empty slice [].
func topoSortBFS(graph map[int][]int, n int) []int {
	if n == 0 {
		return []int{}
	}
	inNode := map[int]int{}
	queue := []int{}

	for node := range graph {
		for _, child := range graph[node] {
			inNode[child] += 1
		}
	}

	for node := 0; node < n; node++ {
		if inNode[node] == 0 {
			queue = append(queue, node)
		}
	}

	index := 0
	for index < len(queue) {
		curr := queue[index]
		index += 1
		for _, child := range graph[curr] {
			inNode[child] -= 1
			if inNode[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	if len(queue) != n {
		return []int{}
	}
	// fmt.Println(inNode)
	// fmt.Println(queue)
	return queue
}

func main() {
	tests := []struct {
		n     int
		edges map[int][]int
	}{
		// Test 1: Multiple dependencies
		{6, map[int][]int{5: {2, 0}, 4: {0, 1}, 2: {3}, 3: {1}}},

		// Test 2: Straight line
		{4, map[int][]int{0: {1}, 1: {2}, 2: {3}}},

		// Test 3: Multiple valid topological orders
		{4, map[int][]int{0: {2}, 2: {3}, 3: {1}}},

		// Test 4: Cycle present
		{3, map[int][]int{0: {1}, 1: {2}, 2: {0}}},

		// Test 5: Larger graph with independent nodes
		{7, map[int][]int{6: {4}, 5: {2, 0}, 4: {0, 1}, 2: {3}, 3: {1}, 1: {}, 0: {}}},

		// Test 6: Empty graph
		{0, map[int][]int{}},

		// Test 7: Graph with isolated nodes only
		{4, map[int][]int{0: {}, 1: {}, 2: {}, 3: {}}},
	}

	for i, t := range tests {
		fmt.Printf("Test %d Output: %v\n", i+1, topoSortBFS(t.edges, t.n))
	}
}
