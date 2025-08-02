// First try - 51 mins 02 secs (own logic)
// Second try - 4 mins 09 secs

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
func rootToNodePath(root *TreeNode, target int) []int {
	res := []int{}
	found := false
	traverse := func(*TreeNode) {}
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		if node.Val == target {
			found = true
		} else {
			traverse(node.Left)
			if found {
				return
			}
			traverse(node.Right)
			if found {
				return
			}
			res = res[:len(res)-1]
		}
	}
	traverse(root)
	return res
}

// // 🚧 Your task: implement this function
// func rootToNodePath(root *TreeNode, target int) []int {
// 	// TODO: implement this
// 	return []int{}
// }

// ✅ Test framework
func runTestCase(name string, root *TreeNode, target int, expected []int) {
	result := rootToNodePath(root, target)
	fmt.Printf("🔍 %s\n", name)
	fmt.Printf("Expected: %v\nGot:      %v\n", expected, result)
	if reflect.DeepEqual(result, expected) {
		fmt.Println("✅ PASS\n")
	} else {
		fmt.Println("❌ FAIL\n")
	}
}

// 🧪 Test cases
func main() {
	// 1. Empty tree
	runTestCase("Empty tree", nil, 1, []int{})

	// 2. Single node match
	single := &TreeNode{Val: 5}
	runTestCase("Single node match", single, 5, []int{5})

	// 3. Single node no match
	runTestCase("Single node no match", single, 10, []int{})

	// 4. Balanced tree
	//       1
	//      / \
	//     2   3
	//    /
	//   4
	tree1 := &TreeNode{1, nil, nil}
	tree1.Left = &TreeNode{2, &TreeNode{4, nil, nil}, nil}
	tree1.Right = &TreeNode{3, nil, nil}
	runTestCase("Target 4 on left", tree1, 4, []int{1, 2, 4})
	runTestCase("Target 3 on right", tree1, 3, []int{1, 3})
	runTestCase("Target not found", tree1, 9, []int{})

	// 5. Deep left-skewed tree
	deep := &TreeNode{1, nil, nil}
	curr := deep
	for i := 2; i <= 10; i++ {
		curr.Left = &TreeNode{i, nil, nil}
		curr = curr.Left
	}
	runTestCase("Deep tree, target 10", deep, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// 6. Deep right-skewed tree
	deepRight := &TreeNode{1, nil, nil}
	curr = deepRight
	for i := 2; i <= 10; i++ {
		curr.Right = &TreeNode{i, nil, nil}
		curr = curr.Right
	}
	runTestCase("Deep right-skewed tree, target 7", deepRight, 7, []int{1, 2, 3, 4, 5, 6, 7})
}
