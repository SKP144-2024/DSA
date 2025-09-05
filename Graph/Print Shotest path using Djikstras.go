package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Vector struct {
	node int
	dist int
}

type MinHeap []Vector

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { (*h) = append(*h, x.(Vector)) }
func (h *MinHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

// Function signature
func dijkstraWithPath(n int, graph map[int][][2]int, src, dest int) ([]int, []int) {
	dist := make([]int, n)
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		parent[i] = -1
	}

	parent[src] = src
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
				parent[val] = curr.node
				heap.Push(h, Vector{val, dist[val]})
			}
		}
	}
	if dist[dest] == math.MaxInt32 {
		return []int{}, []int{}
	}
	node := dest
	path := []int{}
	for parent[node] != node {
		path = append(path, node)
		node = parent[node]
	}
	path = append(path, src)

	d := len(path)
	for i := 0; i < d/2; i++ {
		path[i], path[d-1-i] = path[d-1-i], path[i]
	}

	return dist, path
}

func main() {
	tests := []struct {
		n      int
		graph  map[int][][2]int
		src    int
		dest   int
		output []int
	}{
		// Test 1: Simple linear path
		{
			5,
			map[int][][2]int{
				0: {{1, 1}},
				1: {{2, 1}},
				2: {{3, 1}},
				3: {{4, 1}},
				4: {},
			},
			0,
			4,
			[]int{0, 1, 2, 3, 4},
		},

		// Test 2: Multiple paths, shortest chosen
		{
			4,
			map[int][][2]int{
				0: {{1, 4}, {2, 1}},
				1: {{3, 1}},
				2: {{1, 2}, {3, 5}},
				3: {},
			},
			0,
			3,
			[]int{0, 2, 1, 3},
		},

		// Test 3: Disconnected graph
		{
			4,
			map[int][][2]int{
				0: {{1, 1}},
				1: {},
				2: {{3, 1}},
				3: {},
			},
			0,
			3,
			[]int{},
		},

		// Test 4: Graph with cycle
		{
			5,
			map[int][][2]int{
				0: {{1, 2}, {2, 4}},
				1: {{2, 1}, {3, 7}},
				2: {{4, 3}},
				3: {{4, 1}},
				4: {},
			},
			0,
			4,
			[]int{0, 1, 2, 4},
		},
	}

	for i, test := range tests {
		_, path := dijkstraWithPath(test.n, test.graph, test.src, test.dest)
		if fmt.Sprint(path) == fmt.Sprint(test.output) {
			fmt.Printf("Test %d: PASS ✅\n", i+1)
		} else {
			fmt.Printf("Test %d: FAIL ❌ | got %v, expected %v\n", i+1, path, test.output)
		}
	}
}
