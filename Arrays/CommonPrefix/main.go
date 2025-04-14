package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	min_len := len(strs[0])
	min_word := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < min_len {
			min_len = len(strs[i])
			min_word = strs[i]
		}
	}
	var res []rune
	for ind, letter := range min_word {
		flag := true
		for _, word := range strs {
			if rune(word[ind]) != letter {
				flag = false
				return string(res)
			}

		}
		if flag {
			res = append(res, letter)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{
		"flower", "flow", "flight",
	}))
}
