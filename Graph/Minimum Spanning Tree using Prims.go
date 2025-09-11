package main

import (
	"container/heap"
	"fmt"
)

// ---------- Graph Edge ----------
type Edge struct {
	u, v   int
	weight int
}

type MinHeap []Edge

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Edge)) }
func (h *MinHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

// ---------- TODO: Implement Prim’s Algorithm here ----------
// It should return (totalWeight, mstEdges)
func prims(n int, adj [][]Edge) (int, []Edge) {
	visited := make([]bool, n)
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Edge{-1, 0, 0})

	mst := 0
	mstEdges := []Edge{}
	for pq.Len() > 0 {
		// fmt.Println(pq)
		curr := heap.Pop(pq).(Edge)
		if visited[curr.v] {
			continue
		}
		// fmt.Println("u -> ", curr.u, " v -> ", curr.v, " weight -> ", curr.weight, pq)
		mst += curr.weight
		if curr.u != -1 {
			mstEdges = append(mstEdges, curr)
		}
		visited[curr.v] = true
		for _, edge := range adj[curr.v] {
			if !visited[edge.v] {
				heap.Push(pq, edge)
			}
		}
	}
	if len(mstEdges) < n-1 {
		return 0, []Edge{}
	}
	return mst, mstEdges
}

// ---------- Utility ----------
func makeGraph(n int, edges []Edge) [][]Edge {
	adj := make([][]Edge, n)
	for _, e := range edges {
		adj[e.u] = append(adj[e.u], Edge{e.u, e.v, e.weight})
		adj[e.v] = append(adj[e.v], Edge{e.v, e.u, e.weight})
	}
	return adj
}

func runTest(testName string, n int, edges []Edge, expectedWeight int) {
	adj := makeGraph(n, edges)
	gotWeight, mst := prims(n, adj)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Test:", testName)
	fmt.Println("Graph Edges:", edges)
	fmt.Println("Expected MST Weight:", expectedWeight)
	fmt.Println("Got MST Weight     :", gotWeight)
	fmt.Println("MST Edges Selected :", mst)

	if gotWeight == expectedWeight {
		fmt.Println("✅ PASS")
	} else {
		fmt.Println("❌ FAIL")
	}
	fmt.Println("--------------------------------------------------\n")
}

// ---------- Main ----------
func main() {
	// Testcase 1: Simple 4-node graph
	runTest("Basic Example",
		4,
		[]Edge{
			{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4},
		},
		19,
	)

	// Testcase 2: Triangle graph
	runTest("Triangle Graph",
		3,
		[]Edge{
			{0, 1, 1}, {1, 2, 2}, {0, 2, 3},
		},
		3, // pick edges (0-1,1-2)
	)

	// Testcase 3: Disconnected graph
	runTest("Disconnected Graph",
		4,
		[]Edge{
			{0, 1, 1}, {2, 3, 2},
		},
		0, // Prim will only cover component containing node 0, disconnected graphs are invalid.
	)

	// Testcase 4: Complete graph with equal weights
	runTest("Complete Graph Equal Weights",
		4,
		[]Edge{
			{0, 1, 1}, {0, 2, 1}, {0, 3, 1},
			{1, 2, 1}, {1, 3, 1}, {2, 3, 1},
		},
		3, // any spanning tree with 3 edges of weight 1
	)

	// Testcase 5: Single node
	runTest("Single Node",
		1,
		[]Edge{},
		0, // no edges needed
	)
}
