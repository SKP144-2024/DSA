package main

import (
	"fmt"
	"reflect"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 🚧 Your task: implement this function
func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	// Left bound
	curr := root.Left
	for curr != nil {
		if curr.Left != nil {
			res = append(res, curr.Val)
			curr = curr.Left
		} else if curr.Right != nil {
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			break
		}
	}

	// Leaf Node
	traverse := func(*TreeNode) {}
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && node != root {
			res = append(res, node.Val)
		}
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)

	rbound := []int{}
	// Right bound
	curr = root.Right
	for curr != nil {
		if curr.Right != nil {
			rbound = append(rbound, curr.Val)
			curr = curr.Right
		} else if curr.Left != nil {
			rbound = append(rbound, curr.Val)
			curr = curr.Left
		} else {
			break
		}
	}

	for i := len(rbound) - 1; i >= 0; i-- {
		res = append(res, rbound[i])
	}

	return res
}

// --- Helper functions below ---

// Helper: Build binary tree from level-order input (-1 for nil)
func buildTree(values []int) *TreeNode {
	n := len(values)
	if n == 0 || values[0] == -1 {
		return nil
	}
	nodes := make([]*TreeNode, n)
	for i, v := range values {
		if v != -1 {
			nodes[i] = &TreeNode{Val: v}
		}
	}
	for i := 0; i < n; i++ {
		if nodes[i] == nil {
			continue
		}
		li := 2*i + 1
		ri := 2*i + 2
		if li < n {
			nodes[i].Left = nodes[li]
		}
		if ri < n {
			nodes[i].Right = nodes[ri]
		}
	}
	return nodes[0]
}

// --- Test Framework ---

func runTests() {
	tests := []struct {
		input    []int
		expected []int
		treeArt  string
	}{
		{
			input:    []int{10, 5, 20, 3, 8, 18, 25, -1, -1, 7, -1},
			expected: []int{10, 5, 3, 7, 18, 25, 20},
			treeArt: `
              10
            /    \
           5      20
         /  \    /  \
        3    8  18  25
             /
            7`,
		},
		{
			input:    []int{1},
			expected: []int{1},
			treeArt: `
            1`,
		},
		{
			input:    []int{1, 2, -1, 3, -1, -1, -1, 4},
			expected: []int{1, 2, 3, 4},
			treeArt: `
              1
             /
            2
           /
          3
         /
        4`,
		},
		{
			input:    []int{1, -1, 2, -1, -1, -1, 3, -1, -1, -1, -1, -1, -1, -1, 4},
			expected: []int{1, 4, 3, 2},
			treeArt: `
            1
             \
              2
               \
                3
                 \
                  4`,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 2, 4, 5, 6, 7, 3},
			treeArt: `
              1
            /   \
           2     3
         /  \   / \
        4    5 6   7`,
		},
		{
			input:    []int{1, 2, 3, -1, 4, 5, -1, -1, -1, 6},
			expected: []int{1, 2, 4, 6, 5, 3},
			treeArt: `
              1
            /   \
           2     3
            \   /
             4 5
                /
               6`,
		},
	}

	for i, test := range tests {
		fmt.Printf("\nTest %d:\nBinary Tree:\n%s\n", i+1, test.treeArt)
		root := buildTree(test.input)
		got := boundaryOfBinaryTree(root)
		if reflect.DeepEqual(got, test.expected) {
			fmt.Printf("✅ Passed. Output: %v\n", got)
		} else {
			fmt.Printf("❌ Failed. Expected: %v, Got: %v\n", test.expected, got)
		}
	}
}

// Main
func main() {
	runTests()
}
