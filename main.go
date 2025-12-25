package main

import (
	week4 "300_codes/week_4"
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
	res := week4.SimplifyPath("/home/")
	log.Println("con co", res)
}
