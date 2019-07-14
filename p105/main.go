package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var root *TreeNode
	if len(preorder) != len(inorder) {
		return nil
	}
	doBuildTree(preorder, inorder, 0, len(preorder) - 1, 0, len(inorder) - 1, &root)
	return root
}

func doBuildTree(preorder []int, inorder[]int, pi, pj, ii, ij int, link **TreeNode) {
	if pi > pj || ii > ij {
		return
	}

	rootVal := preorder[pi]
	*link = &TreeNode{Val:rootVal}
	i := ii
	for ; i <= ij; i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	leftCount := i - ii
	pLeftJ := pi + 1 + leftCount - 1
	doBuildTree(preorder, inorder, pi+1, pLeftJ, ii, i - 1, &((*link).Left))
	doBuildTree(preorder, inorder, pLeftJ+1, pj, i+1, ij, &((*link).Right))
}