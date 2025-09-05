package main

import (
	"container/heap"
	"fmt"
)

type Vector struct {
	node int
	dist int
}

type MinHeap []Vector

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Vector)) }
func (h *MinHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val

}

func dijkstra(n int, graph map[int][][2]int, src int) []int {

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = 1e9
	}
	dist[src] = 0
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Vector{src, 0})

	for h.Len() > 0 {
		curr := heap.Pop(h).(Vector)
		if curr.dist > dist[curr.node] {
			continue
		}
		for _, child := range graph[curr.node] {
			val, weight := child[0], child[1]
			if dist[val] > dist[curr.node]+weight {
				dist[val] = dist[curr.node] + weight
				heap.Push(h, Vector{val, dist[val]})
			}
		}
	}

	return dist
}

func main() {
	tests := []struct {
		n      int
		graph  map[int][][2]int
		src    int
		output []int
	}{
		// Test 1: Simple graph
		{
			4,
			map[int][][2]int{
				0: {{1, 4}, {2, 1}},
				1: {{3, 1}},
				2: {{1, 2}, {3, 5}},
				3: {},
			},
			0,
			[]int{0, 3, 1, 4}, // Expected
		},
		// Test 2: Disconnected node
		{
			5,
			map[int][][2]int{
				0: {{1, 2}},
				1: {{2, 3}},
				2: {{3, 1}},
				3: {},
				4: {}, // isolated node
			},
			0,
			[]int{0, 2, 5, 6, 1e9}, // 1e9 means unreachable
		},
		// Test 3: Single node graph
		{
			1,
			map[int][][2]int{
				0: {},
			},
			0,
			[]int{0},
		},
		// Test 4: Graph with multiple paths
		{
			5,
			map[int][][2]int{
				0: {{1, 10}, {2, 3}},
				1: {{2, 1}, {3, 2}},
				2: {{1, 4}, {3, 8}, {4, 2}},
				3: {{4, 7}},
				4: {{3, 9}},
			},
			0,
			[]int{0, 7, 3, 9, 5},
		},
		// Test 5: Larger graph with cycles
		{
			6,
			map[int][][2]int{
				0: {{1, 1}, {2, 5}},
				1: {{2, 2}, {3, 2}},
				2: {{3, 3}, {4, 1}},
				3: {{5, 1}},
				4: {{5, 2}},
				5: {},
			},
			0,
			[]int{0, 1, 3, 3, 4, 4},
		},
		// Test 6: Graph with redundant heap pushes (shows need for visited/stale check)
		{
			3,
			map[int][][2]int{
				0: {{1, 10}, {2, 1}},
				1: {},
				2: {{1, 1}},
			},
			0,
			[]int{0, 2, 1},
		},
	}

	for i, test := range tests {
		result := dijkstra(test.n, test.graph, test.src)
		pass := true
		if len(result) != len(test.output) {
			pass = false
		} else {
			for j := range result {
				if result[j] != test.output[j] {
					pass = false
					break
				}
			}
		}
		if pass {
			fmt.Printf("Test %d: PASS ✅\n", i+1)
		} else {
			fmt.Printf("Test %d: FAIL ❌\n   got: %v\n   expected: %v\n", i+1, result, test.output)
		}
	}
}
