package week5

// 104. Maximum Depth of Binary Tree https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
func MaxDepth(root *TreeNode) int {
	// RECURSIVE WAY
	// var depth func(node *TreeNode) int
	// depth = func(node *TreeNode) int {
	// 	if node == nil {
	// 		return 0
	// 	}
	// 	return 1 + max(depth(node.Left), depth(node.Right))
	// }
	// return depth(root)
	// ITERATIVE WAY (DFS)
	if root == nil {
		return 0
	}
	type NodeDepth struct {
	}
	stack := []*TreeNode{root}
	res := 0
	for len(stack) > 0 {
		n := len(stack)
		for i := 0; i < n; i++ {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node == nil {
				continue
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		res++
	}
	return res
}
