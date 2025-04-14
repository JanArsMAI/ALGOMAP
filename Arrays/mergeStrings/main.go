package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	final := ""
	for i := 0; i != max(len(word1), len(word2)); i++ {
		if i < len(word1) {
			final += string(word1[i])
		}
		if i < len(word2) {
			final += string(word2[i])
		}
	}
	return final
}

func main() {
	test1 := "abcd"
	test2 := "kfg"
	fmt.Println(mergeAlternately(test1, test2))
}
