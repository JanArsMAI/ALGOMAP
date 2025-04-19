package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//можно сделать снова на map, но это решение такое же, как в предыдущей задаче получится, поэтому сделаем на массиве, всё равно гарантирован англ алфавит
	resList := make([]int, 26)
	for i := 0; i < len(s); i++ {
		resList[s[i]-'a']++
		resList[t[i]-'a']--
	}
	for _, val := range resList {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
