package main

import (
	"fmt"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: Implement this function
func kthLargest(root *TreeNode, k int) int {
	count := 0
	number := root.Val
	findLarge := func(*TreeNode) {}
	findLarge = func(node *TreeNode) {
		if node == nil {
			return
		}
		findLarge(node.Right)
		count += 1
		if count == k {
			number = node.Val
			return
		}
		findLarge(node.Left)
	}
	findLarge(root)
	return number // placeholder
}

// Helper: Insert into BST
func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertBST(root.Left, val)
	} else {
		root.Right = insertBST(root.Right, val)
	}
	return root
}

// Helper: Build BST from slice
func buildBST(nums []int) *TreeNode {
	var root *TreeNode
	for _, v := range nums {
		root = insertBST(root, v)
	}
	return root
}

// Test function
func testKthLargest() {
	type testCase struct {
		nums []int
		k    int
		want int
		desc string
	}

	tests := []testCase{
		// 1️⃣ Single node tree
		{nums: []int{5}, k: 1, want: 5, desc: "Single element, k=1"},

		// 2️⃣ Simple right skew
		{nums: []int{2, 3, 4}, k: 1, want: 4, desc: "Right skew tree, largest"},
		{nums: []int{2, 3, 4}, k: 3, want: 2, desc: "Right skew tree, smallest"},

		// 3️⃣ Balanced BST
		{nums: []int{5, 3, 8, 2, 4, 7, 9}, k: 2, want: 8, desc: "Balanced BST, k=2"},
		{nums: []int{5, 3, 8, 2, 4, 7, 9}, k: 5, want: 4, desc: "Balanced BST, k=5"},

		// 4️⃣ Left skew
		{nums: []int{9, 8, 7, 6}, k: 1, want: 9, desc: "Left skew tree, k=1"},
		{nums: []int{9, 8, 7, 6}, k: 4, want: 6, desc: "Left skew tree, k=4"},

		// 5️⃣ Complex BST
		{
			nums: []int{15, 10, 20, 8, 12, 16, 25, 19, 30},
			k:    3, want: 20,
			desc: "Complex BST, k=3",
		},
		{
			nums: []int{15, 10, 20, 8, 12, 16, 25, 19, 30},
			k:    6, want: 15,
			desc: "Complex BST, k=6",
		},
	}

	for _, tc := range tests {
		root := buildBST(tc.nums)
		got := kthLargest(root, tc.k)
		pass := "❌"
		if got == tc.want {
			pass = "✅"
		}
		fmt.Printf("%s | k=%d | Expected=%d, Got=%d | %s\n", pass, tc.k, tc.want, got, tc.desc)
	}
}

func main() {
	testKthLargest()
}
