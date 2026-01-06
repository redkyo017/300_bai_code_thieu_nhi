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
	res := week5.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	log.Println("con co", res)
}
