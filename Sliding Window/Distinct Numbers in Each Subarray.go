package main

import (
	"fmt"
	"reflect"
)

// TODO: Implement this function
func distinctNumbers(nums []int, k int) []int {
	res := []int{}
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] += 1
		if i >= k {
			hash[nums[i-k]] -= 1
			if hash[nums[i-k]] == 0 {
				delete(hash, nums[i-k])
			}
		}
		if i > k-2 {
			res = append(res, len(hash))
		}
	}
	return res
}

// ========== Test Harness ==========
type TestCase struct {
	nums     []int
	k        int
	expected []int
}

func runTests() {
	tests := []TestCase{
		// 1. Example case
		{
			nums:     []int{1, 2, 1, 3, 4, 2, 3},
			k:        4,
			expected: []int{3, 4, 4, 3},
		}, // Visualization:
		// [1,2,1,3] -> {1,2,3} = 3
		// [2,1,3,4] -> {1,2,3,4} = 4
		// [1,3,4,2] -> {1,2,3,4} = 4
		// [3,4,2,3] -> {2,3,4} = 3

		// 2. Small k
		{
			nums:     []int{4, 1, 1},
			k:        2,
			expected: []int{2, 1},
		}, // Visualization:
		// [4,1] -> {4,1} = 2
		// [1,1] -> {1} = 1

		// 3. Increasing sequence
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 3, 3},
		}, // Visualization:
		// [1,2,3] -> {1,2,3} = 3
		// [2,3,4] -> {2,3,4} = 3
		// [3,4,5] -> {3,4,5} = 3

		// 4. All identical numbers
		{
			nums:     []int{7, 7, 7, 7},
			k:        2,
			expected: []int{1, 1, 1},
		}, // Visualization:
		// [7,7] -> {7} = 1
		// [7,7] -> {7} = 1
		// [7,7] -> {7} = 1

		// 5. Window size = 1
		{
			nums:     []int{1, 2, 3},
			k:        1,
			expected: []int{1, 1, 1},
		}, // Visualization:
		// [1] -> {1} = 1
		// [2] -> {2} = 1
		// [3] -> {3} = 1

		// 6. Window size = len(nums)
		{
			nums:     []int{1, 2, 2, 3},
			k:        4,
			expected: []int{3},
		}, // Visualization:
		// [1,2,2,3] -> {1,2,3} = 3

		// 7. Edge case: k > len(nums)
		{
			nums:     []int{1, 2},
			k:        5,
			expected: []int{},
		},
	}

	for i, tc := range tests {
		result := distinctNumbers(tc.nums, tc.k)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("✅ Test %d passed\n", i+1)
		} else {
			fmt.Printf("❌ Test %d failed: got %v, expected %v\n", i+1, result, tc.expected)
		}
	}
}

func main() {
	runTests()
}
