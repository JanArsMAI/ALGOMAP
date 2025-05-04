package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var dummy = &ListNode{}
	dummy.Next = head
	length := 0
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	ans := length - n
	newCur := dummy
	newLen := 0
	for newLen != ans {
		newCur = newCur.Next
		newLen++
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
