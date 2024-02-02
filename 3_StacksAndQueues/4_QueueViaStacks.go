package main

import "fmt"

/*
Queue via Stacks
Implement a MyQueue class which implements a queue using two stacks.
*/
type MyQueue struct {
	s1, s2 *Stack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{s1: &Stack{}, s2: &Stack{}}
}

func (q *MyQueue) Enqueue(v int) {
	q.s1.Push(v)
}

func (q *MyQueue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	q.fill()
	return q.s2.Pop()
}

func (q MyQueue) IsEmpty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}

func (q *MyQueue) fill() {
	if !q.s2.IsEmpty() {
		return
	}
	for !q.s1.IsEmpty() {
		q.s2.Push(q.s1.Pop())
	}
}

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

func (s Stack) IsEmpty() bool {
	return len(s.values) == 0
} 

func main() {
	q := NewMyQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}