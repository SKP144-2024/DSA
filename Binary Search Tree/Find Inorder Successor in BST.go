package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if val < root.Val {
		return findNode(root.Left, val)
	}
	return findNode(root.Right, val)
}

// TODO: Implement this
func inorderSuccessor(root, p *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var suc *TreeNode
	node := root
	for node != nil {
		if node.Val > p.Val {
			if (suc == nil) || (suc.Val > node.Val) {
				suc = node
			}
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return suc
}

func main() {
	tests := []struct {
		values  []int
		target  int
		want    *int // nil means expected successor is nil
		diagram string
	}{
		{
			values: []int{50, 30, 70, 20, 40, 60, 80, 35, 45, 42},
			target: 40,
			want:   ptr(42),
			diagram: `
        50
       /  \
     30    70
    /  \   / \
  20   40 60 80
      /  \
    35    45
          /
        42`,
		},
		{
			values: []int{50, 30, 70, 20, 40, 60, 80, 35, 45, 42},
			target: 45,
			want:   ptr(50),
			diagram: `
        50
       /  \
     30    70
    /  \   / \
  20   40 60 80
      /  \
    35    45
          /
        42`,
		},
		{
			values: []int{50, 30, 70, 20, 40, 60, 80, 72},
			target: 70,
			want:   ptr(72),
			diagram: `
        50
       /  \
     30    70
    /  \   / \
  20   40 60 80
           \
           72`,
		},
		{
			values: []int{50, 30, 70, 20, 40, 60, 80, 75, 72},
			target: 75,
			want:   ptr(80),
			diagram: `
        50
       /  \
     30    70
    /  \   / \
  20   40 60 80
           /
         75
        /
      72`,
		},
		{
			values: []int{15, 10, 20, 8, 12, 16, 25},
			target: 8,
			want:   ptr(10),
			diagram: `
      15
     /  \
   10    20
  / \    / \
 8  12  16 25`,
		},
		{
			values: []int{15, 10, 20, 8, 12, 16, 25},
			target: 25,
			want:   nil,
			diagram: `
      15
     /  \
   10    20
  / \    / \
 8  12  16 25`,
		},
		{
			values: []int{100, 50, 150, 25, 75, 125, 175, 60, 80, 130, 160, 180, 78},
			target: 80,
			want:   ptr(100),
			diagram: `
        100
       /   \
     50     150
    /  \    /  \
  25   75  125 175
       / \   \   / \
     60  80 130 160 180
         /
       78`,
		},
		{
			values: []int{20, 10, 30, 5, 15},
			target: 15,
			want:   ptr(20),
			diagram: `
     20
    /  \
  10    30
  / \
 5  15`,
		},
	}

	for i, tc := range tests {
		var root *TreeNode
		for _, v := range tc.values {
			root = insertBST(root, v)
		}
		node := findNode(root, tc.target)
		got := inorderSuccessor(root, node)
		pass := false
		if tc.want == nil && got == nil {
			pass = true
		} else if tc.want != nil && got != nil && got.Val == *tc.want {
			pass = true
		}

		status := "❌"
		if pass {
			status = "✅"
		}

		fmt.Printf("Test %d %s\nTree:\n%s\nTarget: %d, Expected: %v, Got: %v\n\n",
			i+1, status, tc.diagram, tc.target,
			intPtrVal(tc.want), intPtrVal(valPtr(got)))
	}
}

func ptr(v int) *int { return &v }
func valPtr(t *TreeNode) *int {
	if t == nil {
		return nil
	}
	return &t.Val
}
func intPtrVal(p *int) interface{} {
	if p == nil {
		return nil
	}
	return *p
}
