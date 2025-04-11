package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lMax := math.MinInt
		var lTree []*TreeNode
		for _, node := range queue {
			if node != nil {
				lMax = max(lMax, node.Val)
				if node.Left != nil {
					lTree = append(lTree, node.Left)
				}
				if node.Right != nil {
					lTree = append(lTree, node.Right)
				}
			}
		}
		queue = lTree
		if lMax != math.MinInt {
			res = append(res, lMax)
		}
	}
	return res
}
