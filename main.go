package main

import (
	week3 "300_codes/week_3"
	"log"
)

func main() {
	log.Println("con co be be")
	// nums := []int{-1, 4, 2, 1, 9, 10}
	// nums := []int{-1, -2, -60, 40, 43}
	node1 := week3.ListNode{1, nil}
	node2 := week3.ListNode{2, nil}
	// node3 := week3.ListNode{3, nil}
	// node4 := week3.ListNode{4, nil}
	// node5 := week3.ListNode{5, nil}
	node1.Next = &node2
	// node2.Next = &node3
	// node3.Next = &node4
	// node4.Next = &node5
	res := week3.RemoveNthFromEnd(&node1, 2)
	log.Println("res", res)
}
