package main

import (
	"fmt"
)

func numJewelsInStones(jewels string, stones string) int {
	ansmap := make(map[string]bool)
	counter := 0
	for _, r := range jewels {
		ansmap[string(r)] = true
	}
	for _, r := range stones {
		if _, exists := ansmap[string(r)]; exists {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
