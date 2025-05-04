package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	n := length/2 + 1
	newCur := head
	ans := 1
	for ans != n {
		newCur = newCur.Next
		ans++
	}
	return newCur

}
