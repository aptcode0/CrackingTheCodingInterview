package main

import "fmt"

/*
Stack Min
How would you design a stack which, in addition to push and pop, has a function 
min which returns the minimum eiement? Push, pop and min should all operate in 
O(1) time.
*/
type MinStack struct {
	stack []int
	mins  []int
}

func NewMinStack() *MinStack {
	return &MinStack{}
}

func (s *MinStack) Push(v int) {
	if len(s.mins) == 0 || v <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, v)
	}
	s.stack = append(s.stack, v)
}

func (s *MinStack) Pop() int {
	if len(s.stack) == 0 {
		return -1
	}
	v := s.stack[len(s.stack)-1]
	if s.mins[len(s.mins)-1] == v {
		s.mins = s.mins[:len(s.mins)-1]
	}
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func (s *MinStack) getMin() int {
	if len(s.mins) == 0 {
		return -1
	}
	return s.mins[len(s.mins)-1]
}

func main() {
	ms := NewMinStack()
	ms.Push(3)
	fmt.Println(ms.getMin())
	ms.Push(4)
	fmt.Println(ms.getMin())
	ms.Push(2)
	fmt.Println(ms.getMin())

	ms.Pop()
	fmt.Println(ms.getMin())
}