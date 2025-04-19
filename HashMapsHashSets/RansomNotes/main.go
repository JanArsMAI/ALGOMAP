package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, r := range magazine {
		magazineMap[r]++
	}
	for _, r := range ransomNote {
		res, exists := magazineMap[r]
		if !exists {
			return false
		}
		if res == 0 {
			return false
		}
		magazineMap[r]--
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
