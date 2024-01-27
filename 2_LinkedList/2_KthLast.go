package main

import "fmt"

/*
Return Kth to Last
Implement an algorithm to find the kth to last element of a
singly linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLastRecursive(head *ListNode, k int) *ListNode {
	res, k := kthLastRecursive(head, k)
	return res
}

func kthLastRecursive(head *ListNode, k int) (*ListNode, int) {
	if head == nil {
		return head, k
	}
	node, k := kthLastRecursive(head.Next, k)
	k--
	if k == 0 {
		return head, k
	}
	return node, k
}

func kthLastIterative(head *ListNode, k int) *ListNode {
	front := getNthNode(head, k)
	if front == nil {
		return nil
	}

	back := head
	for front != nil {
		front = front.Next
		back = back.Next
	}
	return back
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

func main() {
	fmt.Println(kthToLastRecursive(createList(), 2).Val)
	fmt.Println(kthLastIterative(createList(), 2).Val)
}

func createList() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	return head
}
