package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	prefix := make([]int, len(nums))
	sufix := make([]int, len(nums))
	prefix[0] = nums[0]
	sufix[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i] * prefix[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		sufix[i] = nums[i] * sufix[i+1]
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[i] = sufix[i+1]
		} else if i == len(nums)-1 {
			nums[i] = prefix[i-1]
		} else {
			nums[i] = prefix[i-1] * sufix[i+1]
		}
	}
	return nums
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
