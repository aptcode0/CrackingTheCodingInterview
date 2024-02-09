package main

/*
Check Subtree: T1 and T2 are two very large binary trees, with T1 much bigger than T2. 
Create an
algorithm to determine if T2 is a subtree of T1.
A tree T2 is a subtree of T1 if there exists a node n in T1 such that the subtree of n 
is identical to 12. 
*/

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func checkSubtree(t1, t2 *TreeNode) bool {
	if t1 == nil {
		return false
	}
	if t2 == nil {
		return true
	}
	return isSame(t1, t2) || checkSubtree(t1.Left, t2) || checkSubtree(t1.Right, t2)
}

func isSame(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}

func main() {
	t1 := createMinimalBST([]int{7, 8, 9, 0, 1, 4, 3})
	t2 := createMinimalBST([]int{1, 4, 3})
	fmt.Println(checkSubtree(t1, t2))
	t2 = createMinimalBST([]int{4, 8, 3})
	fmt.Println(checkSubtree(t1, t2))
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
