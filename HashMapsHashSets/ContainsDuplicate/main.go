package main

import "fmt"

func containsDuplicate(nums []int) bool {
	ansmap := make(map[int]int)
	for _, val := range nums {
		ansmap[val]++
		if ansmap[val] >= 2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
