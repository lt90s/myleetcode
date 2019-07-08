package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	for tmp := head; tmp != nil; tmp=tmp.Next {
		length++
	}

	if length < 2 {
		return head
	}

	if k = k % length; k == 0 {
		return head
	}

	k = length - k
	var prev *ListNode
	p := head
	for k > 0 {
		prev = p
		p = p.Next
		k--
	}
	q := p
	for q.Next != nil {
		q = q.Next
	}

	prev.Next = nil
	q.Next = head
	return p
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("empty")
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
	head := &ListNode{Val:1,Next:&ListNode{Val:2,Next:&ListNode{Val:3,Next:&ListNode{Val:4, Next:&ListNode{Val:5}}}}}
	printList(head)
	printList(rotateRight(head, 2))
}