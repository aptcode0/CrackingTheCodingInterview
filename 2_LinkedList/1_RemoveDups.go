package main

import "fmt"

/*
Remove Dups
Write code to remove duplicates from an unsorted linked list. FOLLOW UP
How would you solve this problem if a temporary buffer is not allowed?
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func removeDupsUsingMap(head *ListNode) {
	seen := map[int]bool{}

	var prev *ListNode
	curr := head

	for curr != nil {
		if seen[curr.Val] {
			prev.Next = curr.Next
			curr = curr.Next
			continue
		}
		prev = curr
		seen[curr.Val] = true
		curr = curr.Next
	}
}

func removeDupsNoMap(head *ListNode) {
	curr := head
	for curr != nil {
		runner := curr
		for runner.Next != nil {
			if runner.Next.Val == curr.Val {
				runner.Next = runner.Next.Next
				continue
			}
			runner = runner.Next
		}
		curr = curr.Next
	}
}

func main() {
	head := createList()
	display(head)
	
	removeDupsUsingMap(head)
	display(head)

	head = createList()
	removeDupsNoMap(head)
	display(head)
}

func createList() *ListNode {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 5}
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