package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (it *BSTIterator) push(root *TreeNode) {
	it.s = append(it.s, root)
}

func (it *BSTIterator) pop() *TreeNode {
	if it.empty() {
		return nil
	}
	length := len(it.s)
	node := it.s[length-1]
	it.s = it.s[:length-1]
	return node
}

func (it *BSTIterator) empty() bool {
	return len(it.s) == 0
}

type BSTIterator struct {
	s []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		s: make([]*TreeNode, 0),
	}

	for root != nil {
		it.push(root)
		root = root.Left
	}
	return it
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.empty() {
		return 0
	}

	node := this.pop()
	if node.Right != nil {
		this.push(node.Right)
		tmp := node.Right
		for tmp.Left != nil {
			this.push(tmp.Left)
			tmp = tmp.Left
		}
	}
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return !this.empty()
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	three := TreeNode{Val: 3}
	nine := TreeNode{Val: 9}
	twenty := TreeNode{Val: 20}
	fifteen := TreeNode{Val: 15, Left: &nine, Right: &twenty}
	seven := TreeNode{Val: 7, Left: &three, Right: &fifteen}

	iterator := Constructor(&seven)

	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
}
