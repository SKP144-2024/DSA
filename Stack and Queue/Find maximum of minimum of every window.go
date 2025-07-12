// // first try - 38mins 26secs
// // Second try - 14mins 54secs
// third try - 11mins 03secs

package main

import (
	"fmt"
	"reflect"
)

type pair struct {
	val, index int
}

type Stack struct {
	data []pair
}

func (s *Stack) push(x pair) { s.data = append(s.data, x) }
func (s *Stack) pop() pair {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}
func (s *Stack) top() pair     { return s.data[len(s.data)-1] }
func (s *Stack) isEmpty() bool { return len(s.data) == 0 }
func (s *Stack) assignMaxMin(res *[]int, maxIndex int) {
	pop := s.pop()
	psIndex := -1
	if !s.isEmpty() {
		psIndex = s.top().index
	}
	win := maxIndex - psIndex - 2
	if (*res)[win] < pop.val {
		(*res)[win] = pop.val
	}
}

// 🔒 User needs to implement this
func maxOfMin(arr []int) []int {
	s := Stack{}
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for !s.isEmpty() && s.top().val > arr[i] {
			s.assignMaxMin(&res, i)
		}
		s.push(pair{val: arr[i], index: i})
	}
	for !s.isEmpty() {
		s.assignMaxMin(&res, len(arr))
	}
	min := res[len(res)-1]
	for i := len(res) - 2; i >= 0; i-- {
		if res[i] > min {
			min = res[i]
		} else if res[i] == 0 {
			res[i] = min
		}
	}
	return res
}

// ✅ Helper function to run test cases
func runTestCase(arr []int, expected []int) {
	output := maxOfMin(arr)
	pass := reflect.DeepEqual(output, expected)

	fmt.Println("Input:     ", arr)
	fmt.Println("Expected:  ", expected)
	fmt.Println("Your Output:", output)

	if pass {
		fmt.Println("✅ PASS")
	} else {
		fmt.Println("❌ FAIL")
	}
	fmt.Println("----------------------------------------")
}

// 🧪 All edge & normal test cases
func main() {
	runTestCase(
		[]int{10, 20, 30, 50, 10, 70, 30},
		[]int{70, 30, 20, 10, 10, 10, 10},
	)

	runTestCase(
		[]int{1, 2, 3, 4, 5},
		[]int{5, 4, 3, 2, 1},
	)

	runTestCase(
		[]int{5, 4, 3, 2, 1},
		[]int{5, 4, 3, 2, 1},
	)

	runTestCase(
		[]int{1, 3, 2, 4, 6, 1, 5},
		[]int{6, 4, 2, 2, 1, 1, 1},
	)

	runTestCase(
		[]int{4},
		[]int{4},
	)

	runTestCase(
		[]int{7, 7, 7, 7},
		[]int{7, 7, 7, 7},
	)

	runTestCase(
		[]int{1, 100, 1, 1, 100, 1},
		[]int{100, 1, 1, 1, 1, 1},
	)

	runTestCase(
		[]int{2, 1, 2},
		[]int{2, 1, 1},
	)
}

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // 🔒 User needs to implement this
// func maxOfMin(arr []int) []int {
// 	return []int{}
// }

// // ✅ Helper function to run test cases
// func runTestCase(arr []int, expected []int) {
// 	output := maxOfMin(arr)
// 	pass := reflect.DeepEqual(output, expected)

// 	fmt.Println("Input:     ", arr)
// 	fmt.Println("Expected:  ", expected)
// 	fmt.Println("Your Output:", output)

// 	if pass {
// 		fmt.Println("✅ PASS")
// 	} else {
// 		fmt.Println("❌ FAIL")
// 	}
// 	fmt.Println("----------------------------------------")
// }

// // 🧪 All edge & normal test cases
// func main() {
// 	runTestCase(
// 		[]int{10, 20, 30, 50, 10, 70, 30},
// 		[]int{70, 30, 20, 10, 10, 10, 10},
// 	)

// 	runTestCase(
// 		[]int{1, 2, 3, 4, 5},
// 		[]int{5, 4, 3, 2, 1},
// 	)

// 	runTestCase(
// 		[]int{5, 4, 3, 2, 1},
// 		[]int{5, 4, 3, 2, 1},
// 	)

// 	runTestCase(
// 		[]int{1, 3, 2, 4, 6, 1, 5},
// 		[]int{6, 4, 2, 2, 1, 1, 1},
// 	)

// 	runTestCase(
// 		[]int{4},
// 		[]int{4},
// 	)

// 	runTestCase(
// 		[]int{7, 7, 7, 7},
// 		[]int{7, 7, 7, 7},
// 	)

// 	runTestCase(
// 		[]int{1, 100, 1, 1, 100, 1},
// 		[]int{100, 1, 1, 1, 1, 1},
// 	)

// 	runTestCase(
// 		[]int{2, 1, 2},
// 		[]int{2, 1, 1},
// 	)
// }
