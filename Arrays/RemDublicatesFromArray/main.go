package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	unique := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[unique] {
			unique++
			nums[unique] = nums[i]
		}
	}
	return unique + 1
}

func main() {
	array := []int{1, 1, 2}
	res := removeDuplicates(array)
	fmt.Println(array[:res], res)
}
