package main

import "fmt"

/*
Stack of Plates
Imagine a (literal) stack of plates. If the stack gets too high, 
it might topple. Therefore, in real life, we would likely start a new stack 
when the previous stack exceeds some threshold. 
Implement a data structure SetOfStacks that mimics this.
Implement a function popAt(int index) which performs a pop operation on a 
specific sub-stack.
*/
type Stack struct {
	values []int
	cap int
}


func NewStack(n int) *Stack {
	return &Stack{cap: n}
}

func (s *Stack) Push(v int) {
	if len(s.values) == s.cap {
		return 
	}
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

func (s *Stack) removeBottom() int {
	if len(s.values) == 0 {
		return -1
	}
	v := s.values[0]
	
	s.values = s.values[1:]
	return v
}

func (s *Stack) Remove(top bool) int {
	if top {
		return s.Pop()
	}
	return s.removeBottom()
}

func (s Stack) IsFull() bool {
	return len(s.values) == s.cap
} 

func (s Stack) IsEmpty() bool {
	return len(s.values) == 0
} 

type SetOfStacks struct {
	stacks []*Stack
	cap int
}

func NewSetOfStacks(threshold int) *SetOfStacks {
	return &SetOfStacks{cap: threshold}
}

func (s *SetOfStacks) Push(v int) {
	if len(s.stacks) > 0 && !s.stacks[len(s.stacks)-1].IsFull() {
		s.stacks[len(s.stacks)-1].Push(v)
		return
	}
	ns := NewStack(s.cap)
	ns.Push(v)
	s.stacks = append(s.stacks, ns)
}

func ( s *SetOfStacks) Pop() int {
	if len(s.stacks) == 0 {
		return -1
	}
	last := s.stacks[len(s.stacks)-1]
	v := last.Pop()
	if last.IsEmpty() {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
	return v
} 

func (s SetOfStacks) StacksCount() int {
	return len(s.stacks)
}

func (s *SetOfStacks) PopAt(stackPos int) int {
	if stackPos >= s.StacksCount() {
		return -1
	}

	return s.shift(stackPos, true)
}

func (s *SetOfStacks) shift(sidx int, removeTop bool) int {
	stack := s.stacks[sidx]
	v := stack.Remove(removeTop)
	if stack.IsEmpty() {
		s.removeStack(sidx)
		return v
	}
	if sidx == len(s.stacks)-1 {
		return v
	}
	stack.Push(s.shift(sidx+1, false))
	return v
}

func (s *SetOfStacks) removeStack(sidx int) {
	s.stacks = append(s.stacks[:sidx], s.stacks[sidx+1:]...)
}

func main() {
	st := NewSetOfStacks(2)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st.StacksCount())
	fmt.Println(st.Pop())
	fmt.Println(st.StacksCount())

	//[[1,2], [3,4], [5]]
	st.Push(3)
	st.Push(4)
	st.Push(5)

	fmt.Println(st.StacksCount())
	fmt.Println(st.PopAt(1))
	
	fmt.Println(st.StacksCount())
}