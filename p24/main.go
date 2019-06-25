package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var afterHead, lastNode *ListNode
	p := head
	for {
		if p == nil {
			break
		}
		q := p.Next
		if q == nil {
			if afterHead == nil {
				afterHead = p
			}
			break
		}
		r := q.Next

		q.Next = p
		p.Next = r
		if lastNode != nil {
			lastNode.Next = q
		}
		lastNode = p

		if afterHead == nil {
			afterHead = q
		}

		p = r
	}
	return afterHead
}

func showList(head *ListNode) {
	if head == nil {
		fmt.Println("[]")
	}
	fmt.Printf("%d", head.Val)
	for head.Next != nil {
		fmt.Printf("->%d", head.Next.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	showList(l)
	showList(swapPairs(l))
}
