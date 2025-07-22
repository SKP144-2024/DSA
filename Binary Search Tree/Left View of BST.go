package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 🔒 Implement this function
// func leftView(root *TreeNode) []int {
// 	leftView := []int{}
// 	depth := 0
// 	maxDep := 0
// 	var traverse func(*TreeNode)
// 	traverse = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		// fmt.Println(depth, node.Val)
// 		depth += 1
// 		if depth > maxDep {
// 			maxDep = depth
// 			leftView = append(leftView, node.Val)
// 		}
// 		traverse(node.Left)
// 		traverse(node.Right)
// 		depth -= 1
// 	}
// 	traverse(root)
// 	return leftView
// }

type Pair struct {
	node  *TreeNode
	level int
}

type Queue struct {
	data []Pair
}

func (q *Queue) enqueue(x Pair) { q.data = append(q.data, x) }
func (q *Queue) dequeue() Pair {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}
func (q *Queue) front() Pair   { return q.data[0] }
func (q *Queue) isEmpty() bool { return len(q.data) == 0 }

// 🔒 Implement this function
func leftView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftview := []int{}
	queue := Queue{}
	seenLevel := map[int]int{}
	queue.enqueue(Pair{root, 0})
	for !queue.isEmpty() {
		curr := queue.dequeue()

		if _, ok := seenLevel[curr.level]; !ok {
			leftview = append(leftview, curr.node.Val)
			seenLevel[curr.level] = curr.node.Val
		}

		if curr.node.Left != nil {
			queue.enqueue(Pair{curr.node.Left, curr.level + 1})
		}
		if curr.node.Right != nil {
			queue.enqueue(Pair{curr.node.Right, curr.level + 1})
		}
	}
	return leftview
}

// ✅ Helper to run test cases
func runTestCase(name string, root *TreeNode, expected []int) {
	output := leftView(root)
	pass := reflect.DeepEqual(output, expected)

	fmt.Printf("Test: %s\n", name)
	fmt.Printf("Expected: %v\n", expected)
	fmt.Printf("Your Output: %v\n", output)
	if pass {
		fmt.Println("✅ PASS")
	} else {
		fmt.Println("❌ FAIL")
	}
	fmt.Println("--------------------------------------------------")
}

// 🧪 All test cases including complex ones
func main() {
	// 1. Basic BST
	root1 := &TreeNode{10,
		&TreeNode{6,
			&TreeNode{4, nil, nil},
			&TreeNode{8, nil, &TreeNode{9, nil, nil}},
		},
		&TreeNode{15, nil, &TreeNode{20, nil, nil}},
	}
	runTestCase("Basic BST", root1, []int{10, 6, 4, 9})

	// 2. Complete Tree
	root2 := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}},
	}
	runTestCase("Complete Tree", root2, []int{1, 2, 4})

	// 3. Skewed Left
	root3 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3,
				&TreeNode{4, nil, nil},
				nil},
			nil},
		nil,
	}
	runTestCase("Skewed Left", root3, []int{1, 2, 3, 4})

	// 4. Skewed Right
	root4 := &TreeNode{1,
		nil,
		&TreeNode{2,
			nil,
			&TreeNode{3,
				nil,
				&TreeNode{4, nil, nil}}},
	}
	runTestCase("Skewed Right", root4, []int{1, 2, 3, 4})

	// 5. Single Node
	root5 := &TreeNode{42, nil, nil}
	runTestCase("Single Node", root5, []int{42})

	// 6. Empty Tree
	runTestCase("Empty Tree", nil, []int{})

	// 7. Duplicate Values
	root6 := &TreeNode{1,
		&TreeNode{1,
			&TreeNode{1, nil, nil},
			nil},
		&TreeNode{1, nil, nil},
	}
	runTestCase("Duplicate Values", root6, []int{1, 1, 1})

	// 8. Zig-Zag Alternating
	root8 := &TreeNode{1,
		nil,
		&TreeNode{2,
			&TreeNode{3,
				nil,
				&TreeNode{4,
					&TreeNode{5, nil, nil},
					nil},
			},
			nil},
	}
	runTestCase("Zig-Zag Alternating", root8, []int{1, 2, 3, 4, 5})

	// 9. Blocked by Right Subtree
	root9 := &TreeNode{1,
		&TreeNode{2,
			nil,
			&TreeNode{4, nil, nil},
		},
		&TreeNode{3,
			nil,
			&TreeNode{5,
				&TreeNode{6, nil, nil},
				nil},
		},
	}
	runTestCase("Blocked by Right Subtree", root9, []int{1, 2, 4, 6})

	// 10. Deep Tree Depth = 1000 (stress test)
	var deepRoot *TreeNode
	for i := 1000; i >= 1; i-- {
		deepRoot = &TreeNode{i, deepRoot, nil}
	}
	expected := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		expected[i] = i + 1
	}
	runTestCase("Deep Tree Depth=1000", deepRoot, expected)

	// 11. Right Root + Left Chain
	root11 := &TreeNode{1,
		nil,
		&TreeNode{2,
			&TreeNode{3,
				&TreeNode{4, nil, nil},
				nil},
			nil},
	}
	runTestCase("Right Root + Left Chain", root11, []int{1, 2, 3, 4})
}
