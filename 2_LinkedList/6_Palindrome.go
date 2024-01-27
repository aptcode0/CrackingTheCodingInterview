package main

import "fmt"

/*
Palindrome
Implement a function to check if a linked list is a palindrome.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindromeRecursive(head *ListNode) bool {
	_, res := isPalindromeHelper(head, length(head))
	return res
}

func isPalindrome(head *ListNode) bool {
	mid := getMid(head)
	mid = reverse(mid)
	return isEqual(head, mid)
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

func isPalindromeHelper(head *ListNode, ln int) (*ListNode, bool) {
	if head == nil || ln == 0 {
		return head, true
	}
	if ln == 1 {
		return head.Next, true
	}
	node, isPali := isPalindromeHelper(head.Next, ln - 2)
	if !isPali || node.Val != head.Val {
		return node, false
	}

	return node.Next, true
}

func getMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}

	return prev
}

func isEqual(h1, h2 *ListNode) bool {
	for h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1 = h1.Next
		h2 = h2.Next
	}
	return true
}

func main() {
	fmt.Println(isPalindromeRecursive(createList([]int{1, 3, 7, 3, 1})))
	fmt.Println(isPalindromeRecursive(createList([]int{1, 3, 7, 1})))

	fmt.Println(isPalindrome(createList([]int{1, 3, 7, 3, 1})))
	fmt.Println(isPalindrome(createList([]int{1, 3, 7, 1})))

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
