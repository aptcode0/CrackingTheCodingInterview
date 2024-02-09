package main

import "fmt"

/*
Minimal Tree
Given a sorted (increasing order) array with unique integer elements, 
write an algo- rithm to create a binary search tree with minimal height.
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func createMinimalBST(arr []int) *TreeNode {
	return createBST(arr, 0, len(arr)-1)
}

func createBST(arr []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{Val: arr[mid]}
	node.Left = createBST(arr, start, mid-1)
	node.Right = createBST(arr, mid+1, end)
	return node
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	root := createMinimalBST(arr)
	inorder(root)
	fmt.Println()
	fmt.Println(height(root))
}

func inorder(node *TreeNode) {
	if node == nil {
		return 
	}
	inorder(node.Left)
	fmt.Print(node.Val, " ")
	inorder(node.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}