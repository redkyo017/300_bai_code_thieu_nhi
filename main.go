package main

import (
	week2 "300_codes/week_2"
	"log"
)

func main() {
	log.Println("con co be be")
	nums := []int{-1, -1, 1}
	res := week2.SubarraySum(nums, 0)
	log.Println("res", res)
}
