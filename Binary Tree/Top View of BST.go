package main

import (
	"fmt"
	"reflect"
)

// TreeNode defines the structure of a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Pair for node and its horizontal distance
type Pair struct {
	node *TreeNode
	hd   int
}

// Queue for BFS
type Queue []Pair

func (q *Queue) enqueue(p Pair) {
	*q = append(*q, p)
}

func (q *Queue) dequeue() Pair {
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

// // 🔒 Implement this function
// func topView(root *TreeNode) []int {
// 	// Your logic here
// 	return nil
// }

// 🔒 Implement this function
func topView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	type pair struct {
		node     *TreeNode
		colLevel int
	}
	minLeft := 0
	maxRight := 0
	seenLevel := map[int]int{}
	queue := []pair{pair{root, 0}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := seenLevel[curr.colLevel]; !ok {
			seenLevel[curr.colLevel] = curr.node.Val
			if minLeft > curr.colLevel {
				minLeft = curr.colLevel
			}
			if maxRight < curr.colLevel {
				maxRight = curr.colLevel
			}
		}

		if curr.node.Left != nil {
			queue = append(queue, pair{curr.node.Left, curr.colLevel - 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, pair{curr.node.Right, curr.colLevel + 1})
		}
	}

	top := []int{}
	for i := minLeft; i <= maxRight; i++ {
		top = append(top, seenLevel[i])
	}

	return top
}

// ✅ Helper to run test case
func runTestCase(name string, root *TreeNode, expected []int) {
	output := topView(root)
	ok := reflect.DeepEqual(output, expected)
	fmt.Printf("Test: %s\n", name)
	fmt.Println("Expected:", expected)
	fmt.Println("Output:  ", output)
	if ok {
		fmt.Println("✅ PASS")
	} else {
		fmt.Println("❌ FAIL")
	}
	fmt.Println("------------------------")
}

// 🧪 Test cases
func main() {
	// 1. Balanced Tree
	//       1
	//      / \
	//     2   3
	root1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	runTestCase("Balanced Tree", root1, []int{2, 1, 3})

	// 2. Left Skewed Tree
	//     1
	//    /
	//   2
	//  /
	// 3
	left := &TreeNode{3, nil, nil}
	left = &TreeNode{2, left, nil}
	left = &TreeNode{1, left, nil}
	runTestCase("Left Skewed", left, []int{3, 2, 1})

	// 3. Right Skewed Tree
	// 1
	//  \
	//   2
	//    \
	//     3
	right := &TreeNode{1, nil, nil}
	right.Right = &TreeNode{2, nil, nil}
	right.Right.Right = &TreeNode{3, nil, nil}
	runTestCase("Right Skewed", right, []int{1, 2, 3})

	// 4. Complete Binary Tree
	//         1
	//       /   \
	//      2     3
	//     / \   / \
	//    4   5 6   7
	tree4 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3,
			&TreeNode{6, nil, nil},
			&TreeNode{7, nil, nil}},
	}
	runTestCase("Complete Tree", tree4, []int{4, 2, 1, 3, 7})

	// 5. Overlapping HDs
	//         1
	//       /   \
	//      2     3
	//       \     \
	//        4     5
	//       /       \
	//      6         7
	tree5 := &TreeNode{1, nil, nil}
	tree5.Left = &TreeNode{2, nil, nil}
	tree5.Right = &TreeNode{3, nil, nil}
	tree5.Left.Right = &TreeNode{4, nil, nil}
	tree5.Right.Right = &TreeNode{5, nil, nil}
	tree5.Left.Right.Left = &TreeNode{6, nil, nil}
	tree5.Right.Right.Right = &TreeNode{7, nil, nil}
	runTestCase("Overlapping HDs", tree5, []int{2, 1, 3, 5, 7})

	// 6. Single Node
	single := &TreeNode{Val: 42}
	runTestCase("Single Node", single, []int{42})

	// 7. Empty Tree
	runTestCase("Empty Tree", nil, []int{})

	// 8. Deep Left Tree
	// 10
	// /
	// 9
	// ...
	// /
	// 1
	var deepLeft *TreeNode
	for i := 10; i >= 1; i-- {
		deepLeft = &TreeNode{i, deepLeft, nil}
	}
	expected := make([]int, 10)
	for i := 0; i < 10; i++ {
		expected[i] = 10 - i
	}
	runTestCase("Deep Left Tree", deepLeft, expected)
}
