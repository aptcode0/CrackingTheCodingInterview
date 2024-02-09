package main

import "fmt"

/*
Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a binary search tree.
You may assume that each node has a link to its parent.
*/

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func inorderSucc(node *TreeNode) *TreeNode {
	if node.Right != nil {
		return leftMost(node.Right)
	}

	parent := node.Parent
	if parent != nil && parent.Right == node {
		node = parent
		parent = parent.Parent
	}
	return node.Parent
}

func leftMost(node *TreeNode) *TreeNode {
	for node != nil && node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {
	root := createBinaryTree()

	fmt.Println(inorderSucc(root).Val)
	fmt.Println(inorderSucc(root.Left).Val)
	fmt.Println(inorderSucc(root.Left.Right).Val)
}

func createBinaryTree() *TreeNode {
	/*
		1
	     / \
	    2   3
	     \
	     4   		
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Parent: root}
	root.Right = &TreeNode{Val: 3, Parent: root}
	root.Left.Right = &TreeNode{Val: 4, Parent: root.Left}
	return root
}
