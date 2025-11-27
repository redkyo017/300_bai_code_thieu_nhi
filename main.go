package main

import (
	week1 "300_codes/week_1"
	"log"
)

func main() {
	log.Println("con co be be")
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := week1.GroupAnagrams(s)
	log.Println("res", res)
}
