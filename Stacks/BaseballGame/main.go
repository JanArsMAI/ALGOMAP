package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	stack := []int{}
	for _, v := range operations {
		a, err := strconv.Atoi(v)
		if err != nil {
			if v == "C" {
				stack = stack[:len(stack)-1]
			} else if v == "D" {
				stack = append(stack, stack[len(stack)-1]*2)
			} else if v == "+" {
				sum := stack[len(stack)-1] + stack[len(stack)-2]
				stack = append(stack, sum)
			}
		} else {
			stack = append(stack, a)
		}
	}
	finalSum := 0
	for _, v := range stack {
		finalSum += v
	}
	return finalSum
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}
