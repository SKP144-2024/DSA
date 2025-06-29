package main

import (
	"container/heap"
	"fmt"
	"reflect"
)

type ele struct {
	val    int
	iIndex int
	jIndex int
}

type MinHeap []ele

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(ele)) }
func (h *MinHeap) Pop() interface{} {
	minVal := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return minVal
}

// Challenge: Merge K Sorted Arrays
// Given K sorted integer arrays, merge them into a single sorted array and return it.

// Implement this function.
func mergeKSortedArrays(arrays [][]int) []int {
	if len(arrays) == 0 {
		return []int{}
	} else if len(arrays) == 1 {
		return arrays[0]
	}

	sortHeap := &MinHeap{}
	heap.Init(sortHeap)
	for i := 0; i < len(arrays); i++ {
		if len(arrays[i]) > 0 {
			heap.Push(sortHeap, ele{val: arrays[i][0], iIndex: i, jIndex: 0})
		}
	}
	ans := []int{}
	for sortHeap.Len() > 0 {
		minVal := heap.Pop(sortHeap).(ele)
		ans = append(ans, minVal.val)
		i := minVal.iIndex
		j := minVal.jIndex
		if j+1 < len(arrays[i]) {
			heap.Push(sortHeap, ele{val: arrays[i][j+1], iIndex: i, jIndex: j + 1})
		}
	}

	return ans
}

func main() {
	type testCase struct {
		arrays   [][]int
		expected []int
		name     string
	}
	tests := []testCase{
		{
			arrays:   [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6, 9}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6, 9},
			name:     "Typical case",
		},
		{
			arrays:   [][]int{},
			expected: []int{},
			name:     "Empty input",
		},
		{
			arrays:   [][]int{{}},
			expected: []int{},
			name:     "Single empty array",
		},
		{
			arrays:   [][]int{{}, {}, {}},
			expected: []int{},
			name:     "All empty arrays",
		},
		{
			arrays:   [][]int{{1, 3, 5}},
			expected: []int{1, 3, 5},
			name:     "Single array",
		},
		{
			arrays:   [][]int{{1, 2, 3}, {}, {4, 5}},
			expected: []int{1, 2, 3, 4, 5},
			name:     "Some empty arrays",
		},
		{
			arrays:   [][]int{{-5, -1, 0}, {-2, 2, 3}, {1, 4, 6}},
			expected: []int{-5, -2, -1, 0, 1, 2, 3, 4, 6},
			name:     "Negative numbers",
		},
		{
			arrays:   [][]int{{1, 1, 1}, {1, 1}, {1}},
			expected: []int{1, 1, 1, 1, 1, 1},
			name:     "All duplicate values",
		},
		{
			arrays:   [][]int{{1, 5, 9}, {2, 6, 8}, {3, 7, 10}},
			expected: []int{1, 2, 3, 5, 6, 7, 8, 9, 10},
			name:     "Interleaved arrays",
		},
		{
			arrays:   [][]int{{100, 200}, {1, 2, 3}},
			expected: []int{1, 2, 3, 100, 200},
			name:     "Different size arrays",
		},
		{
			arrays:   [][]int{{0}},
			expected: []int{0},
			name:     "Single element array",
		},
	}

	passCount := 0
	for i, tc := range tests {
		got := mergeKSortedArrays(tc.arrays)
		if reflect.DeepEqual(got, tc.expected) {
			fmt.Printf("✅ Test %2d - %-25s : PASS\n", i+1, tc.name)
			passCount++
		} else {
			fmt.Printf("❌ Test %2d - %-25s : FAIL\n", i+1, tc.name)
			fmt.Printf("    Expected: %v\n", tc.expected)
			fmt.Printf("    Got     : %v\n", got)
		}
	}
	fmt.Printf("\n%d/%d tests passed\n", passCount, len(tests))
}
