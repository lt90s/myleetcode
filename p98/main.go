package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var prev *int
	return inOrder(root, &prev)
}

func inOrder(node *TreeNode, prev **int) bool {
	if node == nil {
		return true
	}
	if !inOrder(node.Left, prev) {
		return false
	}

	if *prev != nil && **prev >= node.Val {
		return false
	}
	*prev = &node.Val
	return inOrder(node.Right, prev)
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}
	fmt.Println(isValidBST(root))
}
