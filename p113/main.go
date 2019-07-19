package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil && sum == 0 {
		return nil
	}

	current := []int{}
	rv := [][]int{}
	dfs(root, sum, current, &rv)
	return rv
}

func dfs(node *TreeNode, sum int, current []int, rv *[][]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if sum == node.Val {
			current = append(current, node.Val)
			tmp := make([]int, len(current))
			copy(tmp, current)
			*rv = append(*rv, tmp)
		}
		return
	}
	dfs(node.Left, sum - node.Val, append(current, node.Val), rv)
	dfs(node.Right, sum - node.Val, append(current, node.Val), rv)
}

func main() {
	root := &TreeNode{
		Val:5,
		Left:&TreeNode{
			Val:4,
			Left:&TreeNode{
				Val:11,
				Left:&TreeNode{
					Val:7,
				},
				Right:&TreeNode{
					Val:2,
				},
			},
		},
		Right:&TreeNode{
			Val:8,
			Left:&TreeNode{
				Val:13,
			},
			Right:&TreeNode{
				Val:4,
				Left:&TreeNode{
					Val:5,
				},
				Right:&TreeNode{
					Val:1,
				},
			},
		},
	}
	fmt.Println(pathSum(root, 22))
}