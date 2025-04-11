package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		ls := len(queue)
		for i := 0; i != ls; i++ {
			curNode := queue[i]
			var nextNode *Node
			if i+1 < ls {
				nextNode = queue[i+1]
			}
			curNode.Next = nextNode
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}

		}
		queue = queue[ls:]

	}
	return root
}
