package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l <= r {
		if numbers[r]+numbers[l] < target {
			l++
		} else if numbers[r]+numbers[l] > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
