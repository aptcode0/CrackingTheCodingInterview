package main

import "fmt"

/*
Paths with Sum
You are given a binary tree in which each node contains an integer value (which might be positive or negative).
Design an algorithm to count the number of paths that sum to a given value.
The path does not need to start or end at the root or a leaf, but it must go downwards.
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countPathsWithSum(root *TreeNode, t int) int {
	counts := map[int]int{}
	return countPaths(root, t, 0, counts)
}

func countPaths(root *TreeNode, target, currSum int, counts map[int]int) int {
	if root == nil {
		return 0
	}
	currSum += root.Val
	rem := currSum - target
	paths := counts[rem]
	if currSum == target {
		paths++
	}
	counts[currSum]++ 
	paths += countPaths(root.Left, target, currSum, counts)
	paths += countPaths(root.Right, target, currSum, counts)
	counts[currSum]--
	return paths
}

func main() {
	fmt.Println(countPathsWithSum(createTree(), 8))
}

func createTree() *TreeNode {
	/*
		  10
		 /   \
		5     -2
	     /  \
	    3    2
	   /  \   \
	  3    -2  1  	
	
	*/

	root := &TreeNode{10, &TreeNode{Val: 5}, &TreeNode{Val: -2}}
	root.Left.Left = &TreeNode{3, &TreeNode{Val: 3}, &TreeNode{Val: -2}}
	root.Left.Right = &TreeNode{2, nil, &TreeNode{Val: 1}}
	return root
}
