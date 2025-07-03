package main

import (
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return val, true
	}
	return 0, false
}

func (s *Stack) Peek() (int, bool) {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1], true
	}
	return 0, false
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// SortStack sorts the given stack in ascending order (smallest on top) using
// only stack operations and one auxiliary stack. No recursion allowed.
func SortStack(s *Stack) {
	temp := Stack{}
	for !s.IsEmpty() {
		top, _ := s.Pop()
		// fmt.Println("s pop: ", top)
		for {
			if tempPeek, ok := temp.Peek(); ok && tempPeek > top {
				// fmt.Println("temp top: ", tempPeek)
				tempTop, _ := temp.Pop()
				s.Push(tempTop)
			} else {
				break
			}
		}
		temp.Push(top)
		// fmt.Println(s.data, temp.data)
	}
	for !temp.IsEmpty() {
		tempPop, _ := temp.Pop()
		s.Push(tempPop)
	}
	// fmt.Println(s.data)
}

// Helper: returns true if the stack is sorted in ascending order (smallest on top)
func isSortedAscending(s *Stack) bool {
	for i := len(s.data) - 1; i > 0; i-- {
		if s.data[i] > s.data[i-1] {
			return false
		}
	}
	return true
}

// Helper: print stack from top to bottom
func printStack(s *Stack) {
	fmt.Print("[")
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Print(s.data[i])
		if i > 0 {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

// Helper: runs a test, prints pass/fail with sign
func runTest(testName string, input []int, expected []int) {
	s := &Stack{data: make([]int, len(input))}
	copy(s.data, input)
	SortStack(s)
	pass := isSortedAscending(s) && len(s.data) == len(expected)
	for i := range s.data {
		if s.data[i] != expected[i] {
			pass = false
			break
		}
	}
	fmt.Printf("%s: ", testName)
	printStack(s)
	if pass {
		fmt.Println(" ✅ PASS")
	} else {
		fmt.Println(" ❌ FAIL")
		fmt.Printf("   Expected: ")
		printStack(&Stack{data: expected})
		fmt.Println()
	}
}

func main() {
	// Typical case
	runTest("Test 1 - mixed", []int{3, 1, 4, 2}, []int{4, 3, 2, 1})
	// Already sorted
	runTest("Test 2 - already sorted", []int{1, 2, 3, 4}, []int{4, 3, 2, 1})
	// Reverse order
	runTest("Test 3 - reverse order", []int{4, 3, 2, 1}, []int{4, 3, 2, 1})
	// All same element
	runTest("Test 4 - all equal", []int{5, 5, 5, 5}, []int{5, 5, 5, 5})
	// Contains duplicates
	runTest("Test 5 - duplicates", []int{3, 1, 2, 3, 2}, []int{3, 3, 2, 2, 1})
	// Single element
	runTest("Test 6 - single element", []int{7}, []int{7})
	// Empty stack
	runTest("Test 7 - empty", []int{}, []int{})
	// Negative numbers
	runTest("Test 8 - negatives", []int{-1, -3, 0, 2, -2}, []int{2, 0, -1, -2, -3})
	// Large case
	runTest("Test 9 - large", []int{10, 2, 8, 5, 7, 3, 1, 6, 9, 4}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	// Tricky: alternating high/low
	runTest("Test 10 - alternating", []int{1, 100, 2, 99, 3, 98, 4}, []int{100, 99, 98, 4, 3, 2, 1})
}
