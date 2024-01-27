package main

import "fmt"

/*
Intersection
Given two (singly) linked lists, determine if the two lists intersect.
Return the intersecting node. Note that the intersection is defined based
on reference, not value. That is, if the kth node of the first linked list
is the exact same node (by reference) as the j th node of the second linked list,
then they are intersecting.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func findIntersection(h1, h2 *ListNode) *ListNode {
	if h1 == nil || h2 == nil {
		return nil
	}
	t1, l1 := findTailAndSize(h1)
	t2, l2 := findTailAndSize(h2)
	if t1 != t2 {
		return nil
	}

	if l1 < l2 {
		h1, h2 = h2, h1
	}
	h1 = getNthNode(h1, abs(l1-l2))

	for h1 != h2 {
		h1 = h1.Next
		h2 = h2.Next
	}

	return h1
}

func findTailAndSize(h *ListNode) (*ListNode, int) {
	cnt := 1
	for h.Next != nil {
		h = h.Next
		cnt++
	}
	return h, cnt
}

func getNthNode(head *ListNode, k int) *ListNode {
	curr := head 
	for k > 0 {
		if curr == nil {
			return nil
		}
		curr = curr.Next
		k--
	}
	return curr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	/*
	1 - 2 - 3
		    \ 
		     4 - 5
	 	    /   
	     6 - 7	     
	*/
	h1 := createList([]int{1, 2, 3, 4, 5})
	h2 := createList([]int{6, 7})
	h2.Next.Next = h1.Next.Next.Next

	fmt.Println(findIntersection(h1, h2))

	h1 = createList([]int{1, 2})
	h2 = createList([]int{3, 4})
	fmt.Println(findIntersection(h1, h2))
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
