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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p, q, r  *ListNode = head, nil, head
	for p != nil {
		p = p.Next
		if n == 0 {
			q = r
			r = r.Next
		} else {
			n--
		}
	}

	if n == 0 {
		if q == nil {
			head = head.Next
		} else {
			q.Next = r.Next
		}
	}

	return head
}

func showList(head *ListNode) {

	if head == nil {
		fmt.Println("[]")
		return
	}
	fmt.Printf("%d", head.Val)

	for head.Next != nil {
		fmt.Printf("->%d", head.Next.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	showList(head)
	showList(removeNthFromEnd(head, 2))

	head = &ListNode{Val: 1}
	showList(removeNthFromEnd(head, 1))

	head = &ListNode{Val: 1, Next: &ListNode{Val: 2,}}
	showList(removeNthFromEnd(head, 2))
}
