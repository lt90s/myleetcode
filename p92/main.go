package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var reverseHead, mPrev, nNext *ListNode
	if head == nil || head.Next == nil || m == n {
		return head
	}

	p := head
	i := 1
	for ; i < m; i++ {
		mPrev = p
		p = p.Next
	}

	reverseHead = p
	q := p.Next
	for ; i < n; i++ {
		nNext = q.Next
		q.Next = reverseHead
		reverseHead = q
		q = nNext
	}

	if mPrev == nil {
		head = reverseHead
	} else {
		mPrev.Next = reverseHead
	}

	p.Next = nNext
	return head

}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("[]")
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
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	printList(head)
	printList(reverseBetween(head, 1, 5))
}
