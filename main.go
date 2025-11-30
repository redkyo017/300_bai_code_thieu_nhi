package main

import (
	week2 "300_codes/week_2"
	"log"
)

func main() {
	log.Println("con co be be")
	nums := []int{6, 5, 5}
	res := week2.MajorityElement(nums)
	log.Println("res", res)
}
