package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	collection := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		collection[cur] = &Node{
			Val: cur.Val,
		}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		collection[cur].Next = collection[cur.Next]
		collection[cur].Random = collection[cur.Random]
		cur = cur.Next
	}
	return collection[head]
}
