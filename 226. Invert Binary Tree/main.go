package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution using FIFO
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		currentRoot := queue[0]
		currentRoot.Right, currentRoot.Left = currentRoot.Left, currentRoot.Right
		queue = queue[1:]

		if currentRoot.Right != nil {
			queue = append(queue, currentRoot.Right)
		}
		if currentRoot.Left != nil {
			queue = append(queue, currentRoot.Left)
		}
	}
	return root
}
