package main

import "fmt"

/*
Three in One
Describe how you could use a single array to implement three stacks.
*/

type Stack interface {
	Push(v, stackPos int) bool
	Pop(stackPos int) int
	Peek(stackPos int) int
}

type FixedMultiStack struct {
	values []int
	tops []int
	cap int
	nStacks int 
}

func NewFixedMultiStack(cnt, cap int) *FixedMultiStack {
	values := make([]int, cnt * cap)
	tops := make([]int, cnt)
	return &FixedMultiStack{values, tops, cap, cnt}
}

func (s *FixedMultiStack) Push(v, stackPos int) bool {
	if s.isFull(stackPos) {
		return false
	}
	
	s.values[s.getTopPos(stackPos)] = v
	s.tops[stackPos]++
	return true
}

func (s *FixedMultiStack) Pop(stackPos int) int {
	if s.isEmpty(stackPos) {
		return -1
	}
	s.tops[stackPos]--
	return s.values[s.getTopPos(stackPos)]
}

func (s FixedMultiStack) Peek(stackPos int) int {
	if s.isEmpty(stackPos) {
		return -1
	}
	return s.values[s.getTopPos(stackPos)-1]
}

func (s FixedMultiStack) getTopPos(sidx int) int {
	offset := sidx * s.cap
	return offset + s.tops[sidx]
}

func (s FixedMultiStack) isFull(sidx int) bool {
	return s.tops[sidx] == s.cap 
}

func (s FixedMultiStack) isEmpty(sidx int) bool {
	return s.tops[sidx] == 0
}

type FlexiMultiStack struct {
	values []int
	tops []int
	nxt []int
	free int
}

func NewFlexiMultiStack(stackCnt, totCap int) *FlexiMultiStack {
	values := make([]int, totCap)
	nxt := initNxt(totCap)
	tops := initArr(stackCnt, -1)

	return &FlexiMultiStack{values, tops, nxt, 0}
}

func (s *FlexiMultiStack) Push(v, stackPos int) bool {
	if s.isFull() {
		return false
	}
	i := s.free
	s.values[i] = v
	s.free = s.nxt[i]
	s.nxt[i] = s.tops[stackPos]
	s.tops[stackPos] = i
	return true
}

func (s *FlexiMultiStack) Pop(stackPos int) int {
	if s.isEmpty(stackPos) {
		return -1
	}
	top := s.tops[stackPos]
	s.tops[stackPos] = s.nxt[top]
	s.nxt[top] = s.free
	s.free = top
	return s.values[top]
}

func (s FlexiMultiStack) Peek(stackPos int) int {
	if s.isEmpty(stackPos) {
		return -1
	}
	return s.values[s.tops[stackPos]]
}

func (s FlexiMultiStack) isFull() bool {
	return s.free == -1
}

func (s FlexiMultiStack) isEmpty(sidx int) bool {
	return s.tops[sidx] == -1
}


func initArr(n, v int) []int {
	res := make([]int, n)
	for i, _ := range res {
		res[i] = v
	}
	return res
}

func initNxt(n int) []int {
	nxt := make([]int, n)
	for i := 0; i < n-1; i++ {
		nxt[i] = i+1
	}
	nxt[n-1] = -1
	return nxt
}

func main() {
	var ms Stack
	ms = NewFixedMultiStack(3, 2) // 3 stacks each capacity of 2
	fmt.Println("Fixed Multi Stack")
	checkMultiStack(ms)

	ms = NewFlexiMultiStack(3, 6)
	fmt.Println("Flexi Multi Stack")
	checkMultiStack(ms)
}

func checkMultiStack(ms Stack) {
	fmt.Println(ms.Peek(0), ms.Peek(1), ms.Peek(2))

	out := ms.Push(1, 2) 
	fmt.Println(out, ms.Peek(2))
	
	out = ms.Push(2, 1) 
	fmt.Println(out, ms.Peek(1))

	out = ms.Push(3, 2) 
	fmt.Println(out, ms.Peek(2))

	fmt.Println(ms.Push(4, 2), ms.Peek(2))

	v := ms.Pop(2) 
	fmt.Println(v, ms.Peek(2))

	v = ms.Pop(1) 
	fmt.Println(v, ms.Peek(1))

	fmt.Println(ms.Pop(1), ms.Peek(1))

	out = ms.Push(2, 1) // [[], [2], [1]]
	fmt.Println(out, ms.Peek(1))
}