package main

import "fmt"

/* List of Depths
Given a binary tree, design an algorithm which creates a linked list of all the 
nodes at each depth (e.g., if you have a tree with depth D, you'll have D linked 
lists). */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateDepthListsDFS(root *TreeNode) [][]int {
	return createDepthLists(root, 0, nil)	
}

func createDepthLists(root *TreeNode, depth int, lists [][]int) [][]int {
	if root == nil {
		return lists
	}
	if len(lists) == depth {
		lists = append(lists, []int{})
	}
	lists[depth] = append(lists[depth], root.Val)
	lists = createDepthLists(root.Left, depth+1, lists) 
	lists = createDepthLists(root.Right, depth+1, lists) 
	return lists
} 

func CreateDepthListsBFS(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	curr := []*TreeNode{root}
	var res [][]int
	for len(curr) > 0 {
		parents := curr
		curr = []*TreeNode{}
		var vals []int
		for _, p := range parents {
			vals = append(vals, p.Val)
			if p.Left != nil {
				curr = append(curr, p.Left)
			}
			if p.Right != nil {
				curr = append(curr, p.Right)
			}			
		}
		res = append(res, vals)
	}
	return res

}

func main() {
	root := createBinaryTree()
	fmt.Println(CreateDepthListsDFS(root))
	fmt.Println(CreateDepthListsBFS(root))

}

func createBinaryTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	return root
}
