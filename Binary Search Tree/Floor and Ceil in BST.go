package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: Implement these
func findFloor(root *TreeNode, key int) *TreeNode {
	var floor *TreeNode
	node := root
	for node != nil {
		if node.Val <= key {
			floor = node
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return floor
}

func findCeil(root *TreeNode, key int) *TreeNode {
	var ceil *TreeNode
	node := root
	for node != nil {
		if node.Val >= key {
			ceil = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return ceil
}

// Helper to build BST for tests
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

func buildBST(vals []int) *TreeNode {
	var root *TreeNode
	for _, v := range vals {
		root = insertBST(root, v)
	}
	return root
}

// Test runner with ✅ / ❌
func runTests() {
	/*
	 BST Structure:
	         8
	       /   \
	      4    12
	     / \   / \
	    2  6  10 14
	*/
	root := buildBST([]int{8, 4, 12, 2, 6, 10, 14})

	tests := []struct {
		key         int
		expFloorVal int
		expCeilVal  int
	}{
		{5, 4, 6},
		{13, 12, 14},
		{1, -1, 2},   // Floor doesn't exist
		{15, 14, -1}, // Ceil doesn't exist
		{8, 8, 8},    // Exact match
	}

	for _, t := range tests {
		floorNode := findFloor(root, t.key)
		ceilNode := findCeil(root, t.key)

		passFloor := (floorNode == nil && t.expFloorVal == -1) || (floorNode != nil && floorNode.Val == t.expFloorVal)
		passCeil := (ceilNode == nil && t.expCeilVal == -1) || (ceilNode != nil && ceilNode.Val == t.expCeilVal)

		fmt.Printf("Key: %d | Floor: got %v, exp %d %v | Ceil: got %v, exp %d %v\n",
			t.key,
			nodeVal(floorNode), t.expFloorVal, passSymbol(passFloor),
			nodeVal(ceilNode), t.expCeilVal, passSymbol(passCeil))
	}
}

func nodeVal(n *TreeNode) interface{} {
	if n == nil {
		return nil
	}
	return n.Val
}

func passSymbol(ok bool) string {
	if ok {
		return "✅"
	}
	return "❌"
}

func main() {
	runTests()
}
