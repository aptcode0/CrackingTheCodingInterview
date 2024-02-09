package main

import (
	"fmt"
	"math"
)

/*
Validate BST
Implement a function to check if a binary tree is a binary search tree.
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func ValidateBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt, math.MaxInt)
}

func validateBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val > max {
		return false
	}
	return validateBST(root.Left, min, root.Val) && validateBST(root.Right, root.Val, max)
}

func main() {
	fmt.Println(ValidateBST(createMinimalBST([]int{1, 2, 3, 4, 5})))
	fmt.Println(ValidateBST(createMinimalBST([]int{1, 4, 5, 2, 3})))
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
