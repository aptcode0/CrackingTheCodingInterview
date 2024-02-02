package main

import "fmt"

/*
Sort Stack
Write a program to sort a stack such that the smallest items are on the top. 
You can use an additional temporary stack, but you may not copy the elements 
into any other data structure (such as an array).
*/
type Stack struct {
	values []int
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() int {
	if len(s.values) == 0 {
		return -1
	}
	v := s.values[len(s.values)-1]
	
	s.values = s.values[:len(s.values)-1]
	return v
}

func (s Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.values[len(s.values)-1]
}

func (s *Stack) Sort() {
	tmp := &Stack{}

	for !s.IsEmpty() {
		top := s.Pop()
		for !tmp.IsEmpty() && tmp.Peek() > top {
			s.Push(tmp.Pop())				
		}
		tmp.Push(top)
	}
	s.values = tmp.values
}

func (s Stack) IsEmpty() bool {
	return len(s.values) == 0
} 


func main() {
	s := &Stack{}
	vals := []int{4, 2, 1, 3}
	for _, v := range vals {
		s.Push(v)
	}
	
	s.Sort()

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}