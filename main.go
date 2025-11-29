package main

import (
	week1 "300_codes/week_1"
	"log"
)

func main() {
	log.Println("con co be be")
	s := "ADOBECODEBANC"
	t := "ABC"
	res := week1.MinWindow(s, t)
	log.Println("res", res)
}
