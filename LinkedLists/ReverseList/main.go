package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p *ListNode = nil
	c := head
	for c != nil {
		tmp := c.Next
		c.Next = p
		p = c
		c = tmp
	}
	return p
}
