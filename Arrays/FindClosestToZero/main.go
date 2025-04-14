package main

import (
	"fmt"
	"math"
)

func findClosestNumber(nums []int) int {
	var current int = nums[0]
	var largest int = int(math.Abs(float64(nums[0])))
	for _, val := range nums {
		if int(math.Abs(float64(val))) <= largest {
			if int(math.Abs(float64(val))) == largest {
				if val > current {
					current = val
				}
			} else if int(math.Abs(float64(val))) < largest {
				current = val
				largest = int(math.Abs(float64(val)))
			}
		}
	}
	return current
}

func main() {
	a := []int{-4, -2, 1, 4, 8}
	fmt.Println(findClosestNumber(a))
}
