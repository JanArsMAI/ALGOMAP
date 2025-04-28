package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	} else if len(nums) > 1 {
		n := len(nums)
		l, r := 0, n-1
		res := make([]int, n)
		idx := n - 1
		for l <= r {
			leftSq := nums[l] * nums[l]
			rightSq := nums[r] * nums[r]
			if leftSq > rightSq {
				res[idx] = leftSq
				l++
			} else {
				res[idx] = rightSq
				r--
			}
			idx--
		}

		return res
	}
	return []int{}
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
