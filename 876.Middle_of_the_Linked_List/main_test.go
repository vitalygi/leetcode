package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	curNode := head
	midNode := head
	for curNode != nil && curNode.Next != nil {
		curNode = curNode.Next.Next
		midNode = midNode.Next
	}
	return midNode
}
