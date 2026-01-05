package week5

// 114. Flatten Binary Tree to Linked List https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
func Flatten(root *TreeNode) {
	// DFS
	if root == nil {
		return
	}
	// var dfs func(node *TreeNode) *TreeNode
	// dfs = func(node *TreeNode) *TreeNode {
	// 	if node == nil {
	// 		return nil
	// 	}
	// 	if node.Left == nil && node.Right == nil {
	// 		return node
	// 	}

	// 	leftTail := dfs(node.Left)
	// 	rightTail := dfs(node.Right)

	// 	if leftTail != nil {
	// 		leftTail.Right = node.Right
	// 		node.Right = node.Left
	// 		node.Left = nil
	// 	}

	// 	if rightTail != nil {
	// 		return rightTail
	// 	}
	// 	return leftTail
	// }
	// Google styled
	// var prev *TreeNode
	// var dfs func(node *TreeNode)
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	dfs(node.Right)
	// 	dfs(node.Left)
	// 	node.Right = prev
	// 	node.Left = nil
	// 	prev = node

	// }
	// dfs(root)
	// ITERATIVE "MORRIS-LIKE" APPROACH
	curr := root
	for curr != nil {
		if curr.Left != nil {
			pre := curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}

// 116. Populating Next Right Pointers in Each Node https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	// The BFS Way - Use Level Order template - O(n) space
	// For every node that is popped from the queue, if it's not the last node in that level, set node.Next = queue[0].
	// the O(logN) space Way - recursive
	// if root == nil || root.Left == nil || root.Right == nil {
	// 	return root
	// }
	// root.Next = nil
	// root.Left.Next = root.Right
	// var dfs func(parent *Node)
	// dfs = func(parent *Node) {
	// 	if parent.Left == nil || parent.Right == nil {
	// 		return
	// 	}
	// 	if parent.Next != nil {
	// 		parent.Right.Next = parent.Next.Left // cousin
	// 	}
	// 	parent.Left.Next = parent.Right // sibling
	// 	dfs(parent.Left)
	// 	dfs(parent.Right)
	// }
	// dfs(root)
	// return root
	// the O(1) space Way - iterative
	if root == nil {
		return root
	}
	leftMost := root
	for leftMost.Left != nil {
		curr := leftMost
		for curr != nil {
			curr.Left.Next = curr.Right
			if curr.Next != nil {
				curr.Right.Next = curr.Next.Left
			}
			curr = curr.Next
		}
		leftMost = leftMost.Left
	}
	return root
}
