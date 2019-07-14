package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	rv := [][]int{}
	currentLevel, nextLevel := []*TreeNode{}, []*TreeNode{}

	if root == nil {
		return rv
	}

	currentLevel = append(currentLevel, root)
	for len(currentLevel) != 0 {
		tmp := make([]int, len(currentLevel))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = currentLevel[i].Val
		}
		rv = append(rv, tmp)

		for _, node := range currentLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel, nextLevel = nextLevel, currentLevel[:0]
	}
	return rv
}

func main() {
	root := &TreeNode{Val:3,Left:&TreeNode{Val:9}, Right:&TreeNode{Val:20, Left:&TreeNode{Val:15}, Right:&TreeNode{Val:7}}}
	fmt.Println(levelOrder(root))
}
