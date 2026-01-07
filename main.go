package main

import (
	week5 "300_codes/week_5"
	"log"
)

func main() {
	log.Println("con co be be")
	// node1 := week3.ListNode{1, nil}
	// node2 := week3.ListNode{2, nil}
	// node3 := week3.ListNode{3, nil}
	// node4 := week3.ListNode{4, nil}
	// node5 := week3.ListNode{5, nil}
	// node1.Next = &node2
	// node2.Next = &node3
	// node3.Next = &node4
	// node4.Next = &node5
	// res := week4.BackspaceCompare("xywrrmp", "xywrrmu#p")
	node1 := week5.TreeNode{Val: -1}
	node2 := week5.TreeNode{Val: -2}
	node3 := week5.TreeNode{Val: 10}
	node4 := week5.TreeNode{Val: -6}
	node5 := week5.TreeNode{Val: -3}
	node6 := week5.TreeNode{Val: -6}
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node3.Left = &node5
	node3.Right = &node6
	res := week5.MaxPathSum(&node1)
	log.Println("con co", res)
}
