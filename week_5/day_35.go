package week5

// 98. Validate Binary Search Tree https://leetcode.com/problems/validate-binary-search-tree/description/
func IsValidBST(root *TreeNode) bool {
	// The "Range" Strategy
	if root == nil {
		return true
	}
	// var isValid func(node *TreeNode, min int, max int) bool
	// isValid = func(node *TreeNode, min int, max int) bool {
	// 	if node == nil {
	// 		return true
	// 	}
	// 	if node.Val <= min || node.Val >= max {
	// 		return false
	// 	}
	// 	return isValid(node.Left, min, node.Val) && isValid(node.Right, node.Val, max)
	// }
	// return isValid(root, math.MinInt64, math.MaxInt64)
	// The "Inorder" Way
	var prev *int
	var dfs func(curr *TreeNode) bool
	dfs = func(curr *TreeNode) bool {
		if curr == nil {
			return true
		}
		leftValid := dfs(curr.Left)
		if !leftValid {
			return false
		}
		if prev != nil && curr.Val <= *prev {
			return false
		}
		prev = &curr.Val
		return dfs(curr.Right)
	}
	return dfs(root)
}
