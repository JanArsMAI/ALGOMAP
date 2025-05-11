package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		}
	}
	return nums[l]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
