package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	rv := [][]int{}
	if root == nil {
		return rv
	}
	currentLevel, nextLevel := []*TreeNode{root}, []*TreeNode{}
	reverse := false

	for len(currentLevel) != 0 {
		tmp := make([]int, len(currentLevel))
		if reverse {
			for i := len(tmp) - 1; i >= 0; i-- {
				tmp[len(tmp)-1-i] = currentLevel[i].Val
			}
		} else {
			for i := 0; i < len(tmp); i++ {
				tmp[i] = currentLevel[i].Val
			}
		}
		rv = append(rv, tmp)

		for i := 0; i < len(currentLevel); i++ {
			node := currentLevel[i]
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel, nextLevel = nextLevel, currentLevel[:0]
		reverse = !reverse
	}
	return rv
}

func main() {
	root := &TreeNode{Val:1, Left:&TreeNode{Val:2, Left:&TreeNode{Val:4}}, Right: &TreeNode{Val:3, Right:&TreeNode{Val:5}}}
	fmt.Println(zigzagLevelOrder(root))
}
