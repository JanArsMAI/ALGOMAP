package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxStreak := 0

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		}
	}

	return maxStreak
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
