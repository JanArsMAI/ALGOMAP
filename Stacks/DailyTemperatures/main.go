package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	resStack := [][]int{}
	answer := make([]int, len(temperatures))
	for i, val := range temperatures {
		for len(resStack) > 0 && resStack[len(resStack)-1][0] < val {
			stkI := resStack[len(resStack)-1][1]
			resStack = resStack[:len(resStack)-1]
			answer[stkI] = i - stkI
		}
		resStack = append(resStack, []int{val, i})
	}
	return answer
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
