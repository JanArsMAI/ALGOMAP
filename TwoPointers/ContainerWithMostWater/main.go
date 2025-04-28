package main

import (
	"fmt"
)

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		if area > max {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
