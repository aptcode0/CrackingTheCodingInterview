package main

import (
	"fmt"
)

/*
Sum Lists
You have two numbers represented by a linked list, where each node contains
a single digit. The digits are stored in reverse order, such that the 1's digit is at
the head of the list.
Write a function that adds the two numbers and returns the sum
as a linked list.

FOLLOW UP
Suppose the digits are stored in forward order. Repeat the above problem.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func sumListsBackward(l1, l2 *ListNode) *ListNode {
	carry := 0
	pre := &ListNode{}
	curr := pre
	for l1 != nil || l2 != nil || carry > 0 {
		v1, nxt1 := getAndMove(l1)
		v2, nxt2 := getAndMove(l2)
		l1, l2 = nxt1, nxt2

		sum := v1 + v2 + carry
		carry = sum / 10
		sum %= 10

		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
	}
	return pre.Next
}

func sumListsForward(head1, head2 *ListNode) *ListNode {
	len1, len2 := length(head1), length(head2)
	if len1 < len2 {
		head1, head2 = head2, head1 
	}
	head2 = pad(head2, abs(len1-len2))
	stack := pushToStack(head1, head2)
	
	carry := 0
	var res *ListNode
	for len(stack) > 0 {
		head1, head2 = stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		
		sum := head1.Val + head2.Val + carry
		carry = sum / 10
		sum %= 10
		res = &ListNode{sum, res}
	}
	if carry > 0 {
		return &ListNode{carry, res}
	}
	return res
}

func getAndMove(n *ListNode) (int, *ListNode) {
	if n != nil {
		return n.Val, n.Next
	}
	return 0, n
}

func length(head *ListNode) int {
	curr := head
	res := 0
	for curr != nil {
		res++
		curr = curr.Next
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pad(head *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		head = &ListNode{0, head}
	}
	return head
}

func pushToStack(head1, head2 *ListNode) []*ListNode {
	var stack []*ListNode

	for head1 != nil {
		stack = append(stack, head1)
		stack = append(stack, head2)
		head1, head2 = head1.Next, head2.Next
	}
	return stack
}

func main() {
	l1 := createList([]int{1, 7, 1, 1})
	l2 := createList([]int{5, 9, 2})
	display(sumListsBackward(l1, l2))

	l1 = createList([]int{1, 1, 7, 1})
	l2 = createList([]int{2, 9, 5})
	display(sumListsForward(l1, l2))
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

func display(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
	fmt.Println()
}
