package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	currentNode := root
	if currentNode == nil {
		return &TreeNode{val, nil, nil}
	}
	for {
		if currentNode.Val > val {
			if currentNode.Left == nil {
				currentNode.Left = &TreeNode{val, nil, nil}
				break
			} else {
				currentNode = currentNode.Left
			}

		} else if currentNode.Val < val {
			if currentNode.Right == nil {
				currentNode.Right = &TreeNode{val, nil, nil}
				break
			} else {
				currentNode = currentNode.Right
			}
		}
	}
	return root
}
