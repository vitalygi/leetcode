package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	currentNode := root
	for {
		switch {
		case currentNode == nil:
			return nil
		case currentNode.Val == val:
			return currentNode
		case currentNode.Val < val:
			currentNode = currentNode.Right
		case currentNode.Val > val:
			currentNode = currentNode.Left
		}
	}
}
