package main

/*
First Common Ancestor
Design an algorithm and write code to find the first common ancestor of two nodes in a binary 
tree. Avoid storing additional nodes in a data structure. 
*/
import "fmt"

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
}

func commonAncestor(root, p, q *TreeNode) *TreeNode {
	if !cover(root, p) || !cover(root, q) {
		return nil
	}
	return findCommonAncestor(root, p, q)
}

func findCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pLeft := cover(root.Left, p)
	qLeft := cover(root.Left, q)
	if pLeft != qLeft {
		return root
	}
	child := root.Right
	if pLeft {
		child = root.Left
	}
	return findCommonAncestor(child, p, q)
}

func cover(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}
	return cover(root.Left, p) || cover(root.Right, p)
}

func main() {
	root := createBinaryTree()
	fmt.Println(commonAncestor(root, root.Left, root.Right))
	fmt.Println(commonAncestor(root, root.Left.Right, root.Left.Left))
	fmt.Println(commonAncestor(root, root, root.Right))
}

func createBinaryTree() *TreeNode {
	/*
		1
	     / \
	    2   3
	  /   \
	 4     5  		
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	return root
}

