package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelArray := make([]int, 0)
		for range len(queue) {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				levelArray = append(levelArray, node.Val)
				queue = append(queue, node.Left, node.Right)
			}

		}
		if len(levelArray) > 0 {
			result = append(result, levelArray)
		}

	}
	return result
}
