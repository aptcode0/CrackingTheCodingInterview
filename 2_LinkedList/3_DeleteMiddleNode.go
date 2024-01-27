package main

import "fmt"

/*
Delete Middle Node
Implement an algorithm to delete a node in the middle
(i.e., any node but the first and last node, not necessarily the exact middle)
of a singly linked list, given only access to that node.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddleNode(deleteNode *ListNode) {
	deleteNode.Val = deleteNode.Next.Val
	deleteNode.Next = deleteNode.Next.Next
}

func main() {
	head := createList()
	deleteMiddleNode(head.Next.Next)
	display(head)
}

func createList() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
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
