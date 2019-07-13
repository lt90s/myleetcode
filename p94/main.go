package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return nil
	}
	rv := make([]int, 0)

	node := root
	top := -1
	for top >= 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			top++
			node = node.Left
		}
		node = stack[top]
		top--
		stack = stack[:top+1]

		rv = append(rv, node.Val)
		node = node.Right
	}
	return rv
}

// root root.Left root.Right
func preorderTravsersal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return nil
	}
	rv := make([]int, 0)

	node := root
	top := -1
	for top >= 0 || node != nil {
		for node != nil {
			rv = append(rv, node.Val)
			stack = append(stack, node)
			top++
			node = node.Left
		}
		node = stack[top]
		stack = stack[:top]
		top--
		node = node.Right
	}
	return rv
}

// root.Left root.Right root
func postOrderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	top := -1
	rv := []int{}

	stack = append(stack, root)
	top++

	var prev, cur *TreeNode
	for top >= 0 {
		cur = stack[top]

		if prev == nil || prev.Left == cur || prev.Right == cur {
			if cur.Left != nil {
				stack = append(stack, cur.Left)
				top++
			} else if cur.Right != nil {
				stack = append(stack, cur.Right)
				top++
			}
		} else if cur.Left == prev {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				top++
			}
		} else {
			rv = append(rv, cur.Val)
			stack = stack[:top]
			top--
		}
		prev = cur
	}
	return rv
}

func inorderRecursiveTraversal(root *TreeNode) []int {
	rv := make([]int, 0)
	doInorderRecursiveTraversal(root, &rv)
	return rv
}

func doInorderRecursiveTraversal(node *TreeNode, rv *[]int) {
	if node == nil {
		return
	}
	doInorderRecursiveTraversal(node.Left, rv)
	*rv = append(*rv, node.Val)
	doInorderRecursiveTraversal(node.Right, rv)
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(inorderTraversal(root))
	fmt.Println(inorderRecursiveTraversal(root))
	fmt.Println(preorderTravsersal(root))
	fmt.Println(postOrderTraversal(root))
}
