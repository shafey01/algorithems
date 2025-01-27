package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	cur := root

	for cur != nil || len(stack) > 0 {

		for cur != nil {
			stack = append(stack, cur.Left)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, cur.Val)
		cur = cur.Right
	}

	return result
}
