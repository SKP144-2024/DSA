package main

import (
	"container/heap"
	"fmt"
	"reflect"
	"sort"
)

type Pair struct {
	sum int
	i   int
	j   int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].sum > h[j].sum }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MaxHeap) Pop() interface{} {
	maxVal := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return maxVal
}

func maxCombinations(A []int, B []int, K int) []int {
	if len(A) == 0 || len(B) == 0 || K == 0 {
		return []int{}
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	sumHeap := &MaxHeap{}
	visited := make(map[[2]int]bool)
	ans := []int{}
	sumHeap.Push(Pair{A[0] + B[0], 0, 0})
	visited[[2]int{0, 0}] = true

	for K > 0 && sumHeap.Len() > 0 {
		top := heap.Pop(sumHeap).(Pair)
		ans = append(ans, top.sum)
		i := top.i
		j := top.j

		if (i+1 < len(A)) && !visited[[2]int{i + 1, j}] {
			sumHeap.Push(Pair{A[i+1] + B[j], i, j})
			visited[[2]int{i + 1, j}] = true
		}
		if (j+1 < len(B)) && !visited[[2]int{i, j + 1}] {
			sumHeap.Push(Pair{A[i] + B[j+1], i, j + 1})
			visited[[2]int{i, j + 1}] = true
		}
		K -= 1
	}
	return ans
}

func isEqualUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := append([]int{}, a...)
	bCopy := append([]int{}, b...)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	return reflect.DeepEqual(aCopy, bCopy)
}

// Helper function to run individual test cases
func runMaxCombinationsTest(id int, A, B []int, K int, expected []int) {
	result := maxCombinations(A, B, K)
	status := "FAIL ❌"
	if isEqualUnordered(result, expected) {
		status = "PASS ✅"
	}
	fmt.Printf("Test Case %d: %s\n", id, status)
	fmt.Printf("A: %v, B: %v, K: %d\n", A, B, K)
	fmt.Printf("Expected: %v\n", expected)
	fmt.Printf("Got:      %v\n", result)
	fmt.Println("----------")
}

func main() {
	runMaxCombinationsTest(1,
		[]int{1, 4, 2, 3},
		[]int{2, 5, 1, 6},
		4,
		[]int{10, 9, 9, 8},
	)

	runMaxCombinationsTest(2,
		[]int{-1, -2, 0, 1},
		[]int{-2, -1, 0, 1},
		3,
		[]int{2, 1, 1},
	)

	runMaxCombinationsTest(3,
		[]int{5, 5, 5},
		[]int{5, 5, 5},
		5,
		[]int{10, 10, 10, 10, 10},
	)

	runMaxCombinationsTest(4,
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		0,
		[]int{},
	)

	runMaxCombinationsTest(5,
		[]int{1, 2},
		[]int{3, 4},
		10,
		[]int{6, 5, 5, 4},
	)

	runMaxCombinationsTest(6,
		[]int{},
		[]int{},
		3,
		[]int{},
	)
}
