package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 2, 3, 5, 9, 12}, 2))
}
