package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	right := root.Right

	if prev != nil {
		prev.Right = root
	}

	prev = root
	flatten(root.Left)
	flatten(right)
	root.Left = nil
}
