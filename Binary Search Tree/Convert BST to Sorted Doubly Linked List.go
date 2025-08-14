package main

import (
	"fmt"
)

// Node definition
type Node struct {
	Val   int
	Left  *Node // prev
	Right *Node // next
}

// TODO: Implement this function
func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}
	var prev, head *Node
	traverse := func(*Node) {}
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		traverse(node.Left)
		if prev == nil {
			head = node
		} else {
			node.Left = prev
			prev.Right = node
		}
		prev = node
		traverse(node.Right)
	}
	traverse(root)
	prev.Right = head
	head.Left = prev
	return head
}

// Helper: Validate DLL correctness
func validateDLL(head *Node, expectedVals []int) bool {
	if head == nil && len(expectedVals) == 0 {
		return true
	}
	if head == nil || len(expectedVals) == 0 {
		return false
	}

	// Forward traversal
	curr := head
	for i := 0; i < len(expectedVals); i++ {
		if curr.Val != expectedVals[i] {
			return false
		}
		curr = curr.Right
	}
	if curr != head { // Circular check
		return false
	}

	// Backward traversal
	curr = head.Left
	for i := len(expectedVals) - 1; i >= 0; i-- {
		if curr.Val != expectedVals[i] {
			return false
		}
		curr = curr.Left
	}
	return true
}

// Helper: Build BST from sorted array
func buildBST(sorted []int) *Node {
	if len(sorted) == 0 {
		return nil
	}
	mid := len(sorted) / 2
	root := &Node{Val: sorted[mid]}
	root.Left = buildBST(sorted[:mid])
	root.Right = buildBST(sorted[mid+1:])
	return root
}

// Main test runner
func main() {
	tests := []struct {
		name         string
		bstValues    []int
		expectedList []int
	}{
		{
			name:         "Empty Tree",
			bstValues:    []int{},
			expectedList: []int{},
			/*
			   BST:
			   (empty)
			*/
		},
		{
			name:         "Single Node",
			bstValues:    []int{10},
			expectedList: []int{10},
			/*
			   BST:
			   10
			*/
		},
		{
			name:         "Balanced Tree",
			bstValues:    []int{1, 2, 3, 4, 5},
			expectedList: []int{1, 2, 3, 4, 5},
			/*
			   BST:
			       3
			      / \
			     2   5
			    /   /
			   1   4
			*/
		},
		{
			name:         "Left Skewed Tree",
			bstValues:    []int{1, 2, 3},
			expectedList: []int{1, 2, 3},
			/*
			   BST:
			       2
			      / \
			     1   3
			   (not fully skewed because buildBST makes balanced from sorted array)
			*/
		},
		{
			name:         "Right Skewed Tree",
			bstValues:    []int{5, 6, 7},
			expectedList: []int{5, 6, 7},
			/*
			   BST:
			       6
			      / \
			     5   7
			*/
		},
		{
			name:         "Tree With Duplicates",
			bstValues:    []int{1, 1, 2, 3, 3},
			expectedList: []int{1, 1, 2, 3, 3},
			/*
			   BST:
			       2
			      / \
			     1   3
			    /   / \
			   1   3   (nil)
			*/
		},
	}

	for _, tc := range tests {
		fmt.Println("---------------------------------------------------")
		fmt.Println("Test Case:", tc.name)

		// Build BST and visualize
		root := buildBST(tc.bstValues)

		// Call the function under test
		head := treeToDoublyList(root)

		// Validate
		if validateDLL(head, tc.expectedList) {
			fmt.Println("✅ Passed")
		} else {
			fmt.Println("❌ Failed: Expected", tc.expectedList)
		}
	}
}
