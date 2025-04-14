package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	ans := []string{}
	l := 0

	for l < len(nums) {
		r := l
		for r+1 < len(nums) && nums[r+1] == nums[r]+1 {
			r++
		}

		if l == r {
			ans = append(ans, strconv.Itoa(nums[l]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[l], nums[r]))
		}
		l = r + 1
	}

	return ans
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
