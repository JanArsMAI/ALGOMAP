package main

import "fmt"

func majorityElement(nums []int) int {
	n := len(nums) / 2
	resMap := make(map[int]int)
	for _, v := range nums {
		resMap[v]++
		if resMap[v] > n {
			return v
		}
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
