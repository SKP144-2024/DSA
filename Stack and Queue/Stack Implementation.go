package main

import (
	"fmt"
)

// Stack implementation using slice.
type Stack struct {
	data []int
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack. Returns -1 if stack is empty.
func (s *Stack) Pop() int {
	if len(s.data) > 0 {
		l := len(s.data) - 1
		top := s.data[l]
		s.data = s.data[:l]
		return top
	}
	return -1
}

// Top returns the top element without removing it. Returns -1 if stack is empty.
func (s *Stack) Top() int {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	}
	return -1
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	type op struct {
		method string
		value  int
		want   int
	}

	// Tricky and all possible test cases
	testCases := []struct {
		name string
		ops  []op
	}{
		{
			name: "Push and Pop one element",
			ops: []op{
				{"Push", 5, 0},
				{"Pop", 0, 5},
			},
		},
		{
			name: "Pop empty stack",
			ops: []op{
				{"Pop", 0, -1},
			},
		},
		{
			name: "Top empty stack",
			ops: []op{
				{"Top", 0, -1},
			},
		},
		{
			name: "Push multiple, Pop all",
			ops: []op{
				{"Push", 1, 0},
				{"Push", 2, 0},
				{"Push", 3, 0},
				{"Pop", 0, 3},
				{"Pop", 0, 2},
				{"Pop", 0, 1},
				{"Pop", 0, -1},
			},
		},
		{
			name: "Top after pushes",
			ops: []op{
				{"Push", 10, 0},
				{"Push", 20, 0},
				{"Top", 0, 20},
				{"Pop", 0, 20},
				{"Top", 0, 10},
			},
		},
		{
			name: "IsEmpty scenarios",
			ops: []op{
				{"IsEmpty", 0, 1}, // true
				{"Push", 11, 0},
				{"IsEmpty", 0, 0}, // false
				{"Pop", 0, 11},
				{"IsEmpty", 0, 1}, // true
			},
		},
		{
			name: "Alternating Push and Pop",
			ops: []op{
				{"Push", 7, 0},
				{"Pop", 0, 7},
				{"Push", 8, 0},
				{"Top", 0, 8},
				{"Pop", 0, 8},
				{"Pop", 0, -1},
			},
		},
	}

	passCount := 0
	for idx, tc := range testCases {
		s := &Stack{}
		ok := true
		for _, op := range tc.ops {
			switch op.method {
			case "Push":
				s.Push(op.value)
			case "Pop":
				got := s.Pop()
				if got != op.want {
					ok = false
					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    Pop: expected %d, got %d\n", idx+1, tc.name, op.want, got)
				}
			case "Top":
				got := s.Top()
				if got != op.want {
					ok = false
					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    Top: expected %d, got %d\n", idx+1, tc.name, op.want, got)
				}
			case "IsEmpty":
				got := s.IsEmpty()
				expect := op.want == 1
				if got != expect {
					ok = false
					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    IsEmpty: expected %v, got %v\n", idx+1, tc.name, expect, got)
				}
			}
		}
		if ok {
			fmt.Printf("✅ Test %2d - %-30s : PASS\n", idx+1, tc.name)
			passCount++
		}
	}
	fmt.Printf("\n%d/%d tests passed\n", passCount, len(testCases))
}

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // Stack implementation using slice.
// type Stack struct {
// 	data []int
// }

// // Push adds an element to the top of the stack.
// func (s *Stack) Push(x int) {
// 	// TODO: implement this method
// }

// // Pop removes and returns the top element of the stack. Returns -1 if stack is empty.
// func (s *Stack) Pop() int {
// 	// TODO: implement this method
// 	return -1
// }

// // Top returns the top element without removing it. Returns -1 if stack is empty.
// func (s *Stack) Top() int {
// 	// TODO: implement this method
// 	return -1
// }

// // IsEmpty returns true if the stack is empty.
// func (s *Stack) IsEmpty() bool {
// 	// TODO: implement this method
// 	return false
// }

// func main() {
// 	type op struct {
// 		method string
// 		value  int
// 		want   int
// 	}

// 	// Tricky and all possible test cases
// 	testCases := []struct {
// 		name string
// 		ops  []op
// 	}{
// 		{
// 			name: "Push and Pop one element",
// 			ops: []op{
// 				{"Push", 5, 0},
// 				{"Pop", 0, 5},
// 			},
// 		},
// 		{
// 			name: "Pop empty stack",
// 			ops: []op{
// 				{"Pop", 0, -1},
// 			},
// 		},
// 		{
// 			name: "Top empty stack",
// 			ops: []op{
// 				{"Top", 0, -1},
// 			},
// 		},
// 		{
// 			name: "Push multiple, Pop all",
// 			ops: []op{
// 				{"Push", 1, 0},
// 				{"Push", 2, 0},
// 				{"Push", 3, 0},
// 				{"Pop", 0, 3},
// 				{"Pop", 0, 2},
// 				{"Pop", 0, 1},
// 				{"Pop", 0, -1},
// 			},
// 		},
// 		{
// 			name: "Top after pushes",
// 			ops: []op{
// 				{"Push", 10, 0},
// 				{"Push", 20, 0},
// 				{"Top", 0, 20},
// 				{"Pop", 0, 20},
// 				{"Top", 0, 10},
// 			},
// 		},
// 		{
// 			name: "IsEmpty scenarios",
// 			ops: []op{
// 				{"IsEmpty", 0, 1}, // true
// 				{"Push", 11, 0},
// 				{"IsEmpty", 0, 0}, // false
// 				{"Pop", 0, 11},
// 				{"IsEmpty", 0, 1}, // true
// 			},
// 		},
// 		{
// 			name: "Alternating Push and Pop",
// 			ops: []op{
// 				{"Push", 7, 0},
// 				{"Pop", 0, 7},
// 				{"Push", 8, 0},
// 				{"Top", 0, 8},
// 				{"Pop", 0, 8},
// 				{"Pop", 0, -1},
// 			},
// 		},
// 	}

// 	passCount := 0
// 	for idx, tc := range testCases {
// 		s := &Stack{}
// 		ok := true
// 		for _, op := range tc.ops {
// 			switch op.method {
// 			case "Push":
// 				s.Push(op.value)
// 			case "Pop":
// 				got := s.Pop()
// 				if got != op.want {
// 					ok = false
// 					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    Pop: expected %d, got %d\n", idx+1, tc.name, op.want, got)
// 				}
// 			case "Top":
// 				got := s.Top()
// 				if got != op.want {
// 					ok = false
// 					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    Top: expected %d, got %d\n", idx+1, tc.name, op.want, got)
// 				}
// 			case "IsEmpty":
// 				got := s.IsEmpty()
// 				expect := op.want == 1
// 				if got != expect {
// 					ok = false
// 					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    IsEmpty: expected %v, got %v\n", idx+1, tc.name, expect, got)
// 				}
// 			}
// 		}
// 		if ok {
// 			fmt.Printf("✅ Test %2d - %-30s : PASS\n", idx+1, tc.name)
// 			passCount++
// 		}
// 	}
// 	fmt.Printf("\n%d/%d tests passed\n", passCount, len(testCases))
// }
