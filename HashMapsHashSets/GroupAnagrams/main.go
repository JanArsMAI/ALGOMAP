package main

import "fmt"

//мой вариант, который долго работает)
func EqualMaps(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		v2, ex := m2[k]
		if !ex || v2 != v {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	AllMaps := []map[rune]int{}
	res := [][]string{}
	for _, str := range strs {
		curMap := make(map[rune]int)
		for _, r := range str {
			curMap[r]++
		}
		added := false
		for i := 0; i < len(AllMaps); i++ {
			if EqualMaps(AllMaps[i], curMap) {
				res[i] = append(res[i], str)
				added = true
			}
		}
		if added {
			continue
		} else {
			AllMaps = append(AllMaps, curMap)
			res = append(res, []string{str})
		}
	}
	return res
}

//реализация с литкода через хэш таблицу
func hash(s string) int64 {
	var h int64
	for _, r := range s {
		h += int64(r) * int64(r) * int64(r) * int64(r)
	}
	return h
}

func groupAnagrams2(strs []string) [][]string {
	groups := make(map[int64][]string, len(strs))

	for i := 0; i < len(strs); i++ {
		key := hash(strs[i])
		groups[key] = append(groups[key], strs[i])
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
