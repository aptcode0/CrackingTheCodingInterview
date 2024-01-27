package main

import "fmt"

/*
Loop Detection
Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func findLoopStart(h *ListNode) *ListNode {
	fast, slow := h, h

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil	// No Loop found
	}

	slow = h 
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	h := createList([]int{1, 2, 3, 4, 5, 6})
	tail := findTail(h)
	tail.Next = h.Next.Next
	fmt.Println(findLoopStart(h))

	fmt.Println(findLoopStart(createList([]int{1, 2})))
}

func createList(arr []int) *ListNode {
	pre := &ListNode{}
	curr := pre

	for _, v := range arr {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}

	return pre.Next
}

func findTail(h *ListNode) *ListNode {
	for h.Next != nil {
		h = h.Next
	}
	return h
}
