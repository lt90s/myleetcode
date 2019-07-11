package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var header, beHeader *ListNode
	var prev, bePrev *ListNode
	p := head
	for p != nil {
		if p.Val < x {
			if header == nil {
				header = p
			} else {
				prev.Next = p
			}
			prev = p
		} else {
			if beHeader == nil {
				beHeader = p
			} else {
				bePrev.Next = p
			}
			bePrev = p
		}
		p = p.Next
	}
	if prev != nil {
		prev.Next = beHeader
	}

	if bePrev != nil {
		bePrev.Next = nil
	}

	if header == nil {
		return beHeader
	} else {
		return header
	}
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
	header := &ListNode{Val:1, Next:&ListNode{Val:4, Next:&ListNode{Val:3, Next:&ListNode{Val:2, Next:&ListNode{Val:5, Next:&ListNode{Val:2}}}}}}
	printList(header)
	printList(partition(header, 3))
}