package main

import (
	week1 "300_codes/week_1"
	"log"
)

func main() {
	log.Println("con co be be")
	nums := []int{1, 2, 3, 4}
	res := week1.ProductExceptSelf(nums)
	log.Println("res", res)
}
