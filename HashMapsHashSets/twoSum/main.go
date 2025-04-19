package main

import "fmt"

func twoSum(nums []int, target int) []int {
	resMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		del := target - nums[i]
		if _, exists := resMap[del]; exists {
			return []int{resMap[del], i}
		}
		resMap[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
