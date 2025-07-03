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

// SortStackRecursive sorts the stack so that smallest element is on top (using recursion only).
func SortStackRecursive(s *Stack) {
	if s.IsEmpty() {
		return
	}
	x, _ := s.Pop()
	SortStackRecursive(s)
	insertSorted(s, x)
}

func insertSorted(s *Stack, x int) {
	if sPeek, ok := s.Peek(); (!ok) || (sPeek > x) {
		s.Push(x)
		return
	}
	y, _ := s.Pop()
	insertSorted(s, x)
	s.Push(y)
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
	SortStackRecursive(s)
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
	runTest("Test 1 - mixed", []int{3, 1, 4, 2}, []int{4, 3, 2, 1})
	runTest("Test 2 - already sorted", []int{1, 2, 3, 4}, []int{4, 3, 2, 1})
	runTest("Test 3 - reverse order", []int{4, 3, 2, 1}, []int{4, 3, 2, 1})
	runTest("Test 4 - all equal", []int{5, 5, 5, 5}, []int{5, 5, 5, 5})
	runTest("Test 5 - duplicates", []int{3, 1, 2, 3, 2}, []int{3, 3, 2, 2, 1})
	runTest("Test 6 - single element", []int{7}, []int{7})
	runTest("Test 7 - empty", []int{}, []int{})
	runTest("Test 8 - negatives", []int{-1, -3, 0, 2, -2}, []int{2, 0, -1, -2, -3})
	runTest("Test 9 - large", []int{10, 2, 8, 5, 7, 3, 1, 6, 9, 4}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	runTest("Test 10 - alternating", []int{1, 100, 2, 99, 3, 98, 4}, []int{100, 99, 98, 4, 3, 2, 1})
}
