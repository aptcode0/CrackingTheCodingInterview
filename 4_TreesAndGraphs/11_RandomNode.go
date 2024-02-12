package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Random Node: You are implementing a binary tree class from scratch which, in addition to
insert, find, and delete, has a method getRandomNode() which returns a random node from
the tree. All nodes should be equally likely to be chosen.
Design and implement an algorithm for getRandomNode, and explain how you would implement
the rest of the methods.
*/

type Tree struct {
	root *TreeNode
}

func (t *Tree) insert(v int) *TreeNode {
	if t.root == nil {
		t.root = NewTreeNode(v)
		return t.root
	}
	return t.root.insert(v)
}

func (t *Tree) getRandomNode() *TreeNode {
	if t.root == nil {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(t.root.Size)
	return t.root.findIthNode(index)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
	Size int
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{Val: v, Size: 1}
}
func (t *TreeNode) insert(v int) *TreeNode {
	t.Size++
	if v <= t.Val {
		if t.Left == nil {
			t.Left =  NewTreeNode(v)
			return t.Left
		}
		return t.Left.insert(v)
	}
	if t.Right == nil {
		t.Right =  &TreeNode{Val: v}
		return t.Right
	}
	return t.Right.insert(v)
}

func (t *TreeNode) findIthNode(i int) *TreeNode {
	lsize := 0
	
	if t.Left != nil {
		lsize = t.Left.Size
	}
	if i < lsize {
		return t.Left.findIthNode(i)
	}
	if i == lsize {
		return t 
	}
	return t.Right.findIthNode(i - (lsize + 1))
} 

func main() {
	tree := Tree{}
	for _, v := range []int{3, 2, 5, 4, 1, 7, 9, 8} {
		tree.insert(v)
	}	
	for i := 0; i < 5; i++ {
		fmt.Println(tree.getRandomNode())
	}
}

