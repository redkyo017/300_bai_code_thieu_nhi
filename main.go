package main

import (
	week2 "300_codes/week_2"
	"log"
)

func main() {
	log.Println("con co be be")
	nums := []int{-1, 4, 2, 1, 9, 10}
	// nums := []int{-1, -2, -60, 40, 43}
	res := week2.FirstMissingPositive(nums)
	log.Println("res", res)
}
