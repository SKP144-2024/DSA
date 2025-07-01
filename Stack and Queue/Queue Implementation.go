package main

import (
	"fmt"
)

// Queue implementation using slice.
type Queue struct {
	data []int
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
}

// Dequeue removes and returns the front element of the queue. Returns -1 if queue is empty.
func (q *Queue) Dequeue() int {
	if len(q.data) > 0 {
		tail := q.data[0]
		q.data = q.data[1:]
		return tail
	}
	return -1
}

// Front returns the front element without removing it. Returns -1 if queue is empty.
func (q *Queue) Front() int {
	if len(q.data) > 0 {
		return q.data[0]
	}
	return -1
}

// IsEmpty returns true if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func main() {
	type op struct {
		method string
		value  int
		want   int
	}

	testCases := []struct {
		name string
		ops  []op
	}{
		{
			name: "Enqueue and Dequeue one element",
			ops: []op{
				{"Enqueue", 5, 0},
				{"Dequeue", 0, 5},
			},
		},
		{
			name: "Dequeue empty queue",
			ops: []op{
				{"Dequeue", 0, -1},
			},
		},
		{
			name: "Front empty queue",
			ops: []op{
				{"Front", 0, -1},
			},
		},
		{
			name: "Enqueue multiple, Dequeue all",
			ops: []op{
				{"Enqueue", 1, 0},
				{"Enqueue", 2, 0},
				{"Enqueue", 3, 0},
				{"Dequeue", 0, 1},
				{"Dequeue", 0, 2},
				{"Dequeue", 0, 3},
				{"Dequeue", 0, -1},
			},
		},
		{
			name: "Front after enqueues",
			ops: []op{
				{"Enqueue", 10, 0},
				{"Enqueue", 20, 0},
				{"Front", 0, 10},
				{"Dequeue", 0, 10},
				{"Front", 0, 20},
			},
		},
		{
			name: "IsEmpty scenarios",
			ops: []op{
				{"IsEmpty", 0, 1}, // true
				{"Enqueue", 11, 0},
				{"IsEmpty", 0, 0}, // false
				{"Dequeue", 0, 11},
				{"IsEmpty", 0, 1}, // true
			},
		},
		{
			name: "Alternating Enqueue and Dequeue",
			ops: []op{
				{"Enqueue", 7, 0},
				{"Dequeue", 0, 7},
				{"Enqueue", 8, 0},
				{"Front", 0, 8},
				{"Dequeue", 0, 8},
				{"Dequeue", 0, -1},
			},
		},
	}

	passCount := 0
	for idx, tc := range testCases {
		q := &Queue{}
		ok := true
		for _, op := range tc.ops {
			switch op.method {
			case "Enqueue":
				q.Enqueue(op.value)
			case "Dequeue":
				got := q.Dequeue()
				if got != op.want {
					ok = false
					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    Dequeue: expected %d, got %d\n", idx+1, tc.name, op.want, got)
				}
			case "Front":
				got := q.Front()
				if got != op.want {
					ok = false
					fmt.Printf("❌ Test %2d - %-30s : FAIL\n    Front: expected %d, got %d\n", idx+1, tc.name, op.want, got)
				}
			case "IsEmpty":
				got := q.IsEmpty()
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
