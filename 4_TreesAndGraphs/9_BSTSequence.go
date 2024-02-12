package main

import "fmt"

/*
BST Sequences
A binary search tree was created by traversing through an array from left to right and
inserting each element. Given a binary search tree with distinct elements, print all
possible arrays that could have led to this tree.
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findAllSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{nil}
	}

	leftSeqs := findAllSequences(root.Left) 
	rightSeqs := findAllSequences(root.Right)

	var res [][]int
	curr := []int{root.Val}
	for _, lseq := range leftSeqs {
		for _, rseq := range rightSeqs {
			res = append(res, weave(lseq, rseq, curr, nil)...)
		}
	}
	return res
}

func weave(lseq, rseq, curr []int, res [][]int) [][]int {
	if len(lseq) == 0 || len(rseq) == 0 {
		curr = append(curr, lseq...)
		curr = append(curr, rseq...)
		res = append(res, curr)
		return res
	}
	
	res = weave(lseq[1:], rseq, append(curr, lseq[0]), res)
	res = weave(lseq, rseq[1:], append(curr, rseq[0]), res)
	return res
}


func main() {
	root := createBST()
	fmt.Println(findAllSequences(root))
}

func createBST() *TreeNode {
	/*
		3
	    /   \
	   1     4
	    \
	    2
	*/

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	return root
}