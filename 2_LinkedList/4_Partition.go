package main

import "fmt"

/*
Partition
Write code to partition a linked list around a value x, such that all
nodes less than x come before all nodes greater than or equal to x.
If x is contained within the list, the values of x only need to be
after the elements less than x (see below).
The partition element x can appear anywhere in the "right partition";
it does not need to appear between the left and right partitions.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func orderedPartition(head *ListNode, x int) *ListNode {
	pre1, pre2 := &ListNode{}, &ListNode{}
	before, after := pre1, pre2
	curr := head
	for curr != nil {
		if curr.Val < x {
			before.Next = curr
			before = before.Next
		} else {
			after.Next = curr
			after = after.Next
		}	
		curr = curr.Next
	}

	after.Next = nil
	if pre1.Next == nil {
		return pre2.Next
	} 
	
	before.Next = pre2.Next
	return pre1.Next
}

func unorderedPartition(head *ListNode, x int) *ListNode {
	before, after := head, head
	curr := head

	for curr != nil {
		nxt := curr.Next
		if curr.Val < x {
			curr.Next = before
			before = curr
		} else {
			after.Next = curr
			after = after.Next
		}
		curr = nxt
	}
	after.Next = nil

	return before
}

func main() {
	head := createList()
	head = orderedPartition(head, 35)
	display(head)

	head = createList()
	head = unorderedPartition(head, 35)
	display(head)
}

func createList() *ListNode {
	head := &ListNode{Val: 50}
	head.Next = &ListNode{Val: 30}
	head.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next = &ListNode{Val: 40}
	return head
}

func display(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
	fmt.Println()
}

