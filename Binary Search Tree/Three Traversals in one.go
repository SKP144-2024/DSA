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

type StateNode struct {
	node  *TreeNode
	state int // 1: preorder, 2: inorder, 3: postorder
}

type Stack struct {
	data []StateNode
}

func (s *Stack) push(x StateNode) { s.data = append(s.data, x) }
func (s *Stack) pop() StateNode {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}
func (s *Stack) top() StateNode { return s.data[len(s.data)-1] }
func (s *Stack) isEmpty() bool  { return len(s.data) == 0 }

// 🔒 Implement this
func threeTraversals(root *TreeNode) (preorder, inorder, postorder []int) {
	if root == nil {
		return
	}
	stack := Stack{}
	stack.push(StateNode{root, 1})
	for !stack.isEmpty() {
		curr := stack.pop()
		// fmt.Println(stack)
		if curr.state == 1 {
			preorder = append(preorder, curr.node.Val)
			stack.push(StateNode{curr.node, curr.state + 1})
			if curr.node.Left != nil {
				stack.push(StateNode{curr.node.Left, 1})
			}
		} else if curr.state == 2 {
			inorder = append(inorder, curr.node.Val)
			stack.push(StateNode{curr.node, curr.state + 1})
			if curr.node.Right != nil {
				stack.push(StateNode{curr.node.Right, 1})
			}
		} else if curr.state == 3 {
			postorder = append(postorder, curr.node.Val)
		}
	}
	return
}

func normalize(slice []int) []int {
	if slice == nil {
		return []int{}
	}
	return slice
}

// ✅ Helper to compare and print results
func validateTraversal(name string, root *TreeNode, expPre, expIn, expPost []int) {
	pre, in, post := threeTraversals(root)

	fmt.Println("🔍", name)
	fmt.Println("Preorder:  Expected:", expPre, " | Got:", pre)
	fmt.Println("Inorder:   Expected:", expIn, " | Got:", in)
	fmt.Println("Postorder: Expected:", expPost, " | Got:", post)

	if reflect.DeepEqual(normalize(pre), expPre) &&
		reflect.DeepEqual(normalize(in), expIn) &&
		reflect.DeepEqual(normalize(post), expPost) {
		fmt.Println("✅ PASS")
	} else {
		fmt.Println("❌ FAIL")
	}

}

func main() {
	// 1. Balanced Binary Tree
	root1 := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}
	validateTraversal("Balanced Tree",
		root1,
		[]int{1, 2, 4, 5, 3},
		[]int{4, 2, 5, 1, 3},
		[]int{4, 5, 2, 3, 1},
	)

	// 2. Left Skewed
	left := &TreeNode{3, nil, nil}
	left = &TreeNode{2, left, nil}
	left = &TreeNode{1, left, nil}
	validateTraversal("Left Skewed",
		left,
		[]int{1, 2, 3},
		[]int{3, 2, 1},
		[]int{3, 2, 1},
	)

	// 3. Right Skewed
	right := &TreeNode{1, nil, nil}
	right.Right = &TreeNode{2, nil, nil}
	right.Right.Right = &TreeNode{3, nil, nil}
	validateTraversal("Right Skewed",
		right,
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{3, 2, 1},
	)

	// 4. Full Tree Depth 3
	root4 := &TreeNode{10,
		&TreeNode{20, &TreeNode{40, nil, nil}, &TreeNode{50, nil, nil}},
		&TreeNode{30, &TreeNode{60, nil, nil}, &TreeNode{70, nil, nil}},
	}
	validateTraversal("Full Depth-3 Tree",
		root4,
		[]int{10, 20, 40, 50, 30, 60, 70},
		[]int{40, 20, 50, 10, 60, 30, 70},
		[]int{40, 50, 20, 60, 70, 30, 10},
	)

	// 5. Single Node
	one := &TreeNode{42, nil, nil}
	validateTraversal("Single Node",
		one,
		[]int{42},
		[]int{42},
		[]int{42},
	)

	// 6. Empty Tree
	validateTraversal("Empty Tree", nil, []int{}, []int{}, []int{})
}
