package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)
	return root
}

func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}

	printInOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInOrder(root.Right)
}

func main() {
	one := TreeNode{Val: 1}
	two := TreeNode{Val: 2}
	three := TreeNode{Val: 3}
	four := TreeNode{Val: 4}
	six := TreeNode{Val: 6}
	seven := TreeNode{Val: 7}
	nine := TreeNode{Val: 9}

	four.Left = &two
	four.Right = &seven
	two.Left = &one
	two.Right = &three
	seven.Left = &six
	seven.Right = &nine

	newRoot := invertTree(&four)
	printInOrder(newRoot)
}
