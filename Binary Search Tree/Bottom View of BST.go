package main

import (
	"fmt"
	"reflect"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 🔒 Your task: implement this function
// func bottomView(root *TreeNode) []int {
// 	// TODO: Implement this function
// 	return []int{}
// }

// 🔒 Your task: implement this function
func bottomView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	type Pair struct {
		node     *TreeNode
		colLevel int
	}
	queue := []Pair{Pair{root, 0}}
	maxLeft := 0
	maxRight := 0
	seenLevel := map[int]int{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		seenLevel[curr.colLevel] = curr.node.Val
		if curr.colLevel < maxLeft {
			maxLeft = curr.colLevel
		} else if curr.colLevel > maxRight {
			maxRight = curr.colLevel
		}

		if curr.node.Left != nil {
			queue = append(queue, Pair{curr.node.Left, curr.colLevel - 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, Pair{curr.node.Right, curr.colLevel + 1})
		}

	}
	bottom := []int{}
	for i := maxLeft; i <= maxRight; i++ {
		bottom = append(bottom, seenLevel[i])
	}

	return bottom
}

// ✅ Test runner
func runTestCase(name string, root *TreeNode, expected []int) {
	output := bottomView(root)
	pass := reflect.DeepEqual(output, expected)

	fmt.Println("🔹 Test Case:", name)
	fmt.Println("   Expected:", expected)
	fmt.Println("   Output  :", output)

	if pass {
		fmt.Println("✅ PASS")
	} else {
		fmt.Println("❌ FAIL")
	}
	fmt.Println("--------------------------------------------------")
}

func main() {
	// 1. Empty Tree
	runTestCase("Empty Tree", nil, []int{})

	// 2. Single Node
	runTestCase("Single Node", &TreeNode{1, nil, nil}, []int{1})

	// 3. Left Skewed
	left := &TreeNode{4, nil, nil}
	left = &TreeNode{3, left, nil}
	left = &TreeNode{2, left, nil}
	left = &TreeNode{1, left, nil}
	runTestCase("Left Skewed", left, []int{4, 3, 2, 1})

	// 4. Right Skewed
	right := &TreeNode{1, nil, nil}
	right = &TreeNode{2, nil, right}
	right = &TreeNode{3, nil, right}
	right = &TreeNode{4, nil, right}
	runTestCase("Right Skewed", right, []int{4, 3, 2, 1})

	// 5. Balanced Tree
	root := &TreeNode{20, nil, nil}
	root.Left = &TreeNode{8, nil, nil}
	root.Right = &TreeNode{22, nil, nil}
	root.Left.Left = &TreeNode{5, nil, nil}
	root.Left.Right = &TreeNode{3, nil, nil}
	root.Right.Right = &TreeNode{25, nil, nil}
	root.Left.Right.Left = &TreeNode{10, nil, nil}
	root.Left.Right.Right = &TreeNode{14, nil, nil}
	runTestCase("Balanced Tree", root, []int{5, 10, 3, 14, 25})

	// 6. Complete Tree
	full := &TreeNode{1, nil, nil}
	full.Left = &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}
	full.Right = &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}
	runTestCase("Complete Tree", full, []int{4, 2, 6, 3, 7})

	// 7. Overlapping Horizontal Distances
	tree := &TreeNode{1, nil, nil}
	tree.Left = &TreeNode{2, nil, nil}
	tree.Right = &TreeNode{3, nil, nil}
	tree.Left.Right = &TreeNode{4, nil, nil}
	tree.Right.Right = &TreeNode{5, nil, nil}
	tree.Left.Right.Left = &TreeNode{6, nil, nil}
	tree.Right.Right.Right = &TreeNode{7, nil, nil}
	runTestCase("Overlap HDs", tree, []int{6, 4, 3, 5, 7})

	// 8. Deep Right Tree
	deep := &TreeNode{Val: 1}
	curr := deep
	for i := 2; i <= 10; i++ {
		curr.Right = &TreeNode{Val: i}
		curr = curr.Right
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	runTestCase("Deep Right Tree", deep, expected)
}
