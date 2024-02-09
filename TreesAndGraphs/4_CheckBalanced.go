package main

import "fmt"

/*
Check Balanced
Implement a function to check if a binary tree is balanced. For the purposes of this question,
a balanced tree is defined to be a tree such that the heights of the two subtrees of any node
never differ by more than one.
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return checkBalanced(root) != -1
}

func checkBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := checkBalanced(root.Left)
	if lh == -1 {
		return -1
	}
	rh := checkBalanced(root.Right)
	if rh == -1 {
		return -1
	}
	diff := abs(lh - rh)
	if diff > 1 {
		return -1
	}
	return 1 + max(lh, rh)
}

func main() {
	fmt.Println(isBalanced(createBalancedTree()))
	fmt.Println(isBalanced(createUnBalancedTree()))
}

func createBalancedTree() *TreeNode{
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	return root
}

func createUnBalancedTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	return root
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}