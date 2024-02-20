package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirror(root.Left, root.Right)
}

func mirror(L, R *TreeNode) bool {
	if R == nil && L == nil {
		return true
	}
	if R != nil && L != nil {
		return (L.Val == R.Val) && mirror(L.Left, R.Right) && mirror(L.Right, R.Left)
	}
	return false
}

func main() {
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{1, nil, nil}
	node3 := &TreeNode{3, node1, node2}
	res := isSymmetric(node3)
	fmt.Println("Is the tree symmetric? ", res) // Is the tree symmetric?  false
}
