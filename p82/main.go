package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var prev, p, q *ListNode
	if head == nil {
		return nil
	}
	p = head
	head = nil
	for p != nil {
		q = p.Next
		if q == nil {
			if prev == nil || head == nil {
				return p
			} else {
				prev.Next = p
				return head
			}
		}
		if q.Val == p.Val {
			for q != nil {
				if q.Val != p.Val {
					break
				}
				q = q.Next
			}
			p = q
		} else {
			if prev == nil {
				head, prev = p, p
			} else {
				prev.Next = p
				prev = p
			}
			p = q
		}

	}
	if prev != nil {
		prev.Next = nil
	}
	return head
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Print(head.Val)
	for head.Next != nil {
		fmt.Printf("=>%d", head.Next.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}
	printList(deleteDuplicates(l1))
	l2 := &ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:2, Next:&ListNode{Val:2, Next:&ListNode{Val:2}}}}}
	printList(deleteDuplicates(l2))
}
