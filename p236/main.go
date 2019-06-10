package main

import "fmt"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}

func main() {
	seven := TreeNode{Val: 7}
	four := TreeNode{Val: 4}
	two := TreeNode{Val: 2, Left: &seven, Right: &four}
	six := TreeNode{Val: 6}
	five := TreeNode{Val: 5, Left: &six, Right: &two}
	zero := TreeNode{Val: 0}
	eight := TreeNode{Val: 8}
	one := TreeNode{Val: 1, Left: &zero, Right: &eight}
	three := TreeNode{Val: 3, Left: &five, Right: &one}

	fmt.Println(lowestCommonAncestor(&three, &five, &one))
	fmt.Println(lowestCommonAncestor(&three, &five, &four))
	fmt.Println(lowestCommonAncestor(&three, &five, &eight))
}
