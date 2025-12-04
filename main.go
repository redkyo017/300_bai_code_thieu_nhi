package main

import (
	week2 "300_codes/week_2"
	"log"
)

func main() {
	log.Println("con co be be")
	nums := []int{100, 4, 200, 1, 3, 2}
	res := week2.LongestConsecutive(nums)
	log.Println("res", res)
}
