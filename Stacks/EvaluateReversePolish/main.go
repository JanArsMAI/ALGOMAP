package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stackRes := []string{}
	for _, val := range tokens {
		if val == "+" {
			first, _ := strconv.Atoi(stackRes[len(stackRes)-1])
			second, _ := strconv.Atoi(stackRes[len(stackRes)-2])
			stackRes = append(stackRes[:len(stackRes)-2], strconv.Itoa(first+second))
		} else if val == "-" {
			first, _ := strconv.Atoi(stackRes[len(stackRes)-2])
			second, _ := strconv.Atoi(stackRes[len(stackRes)-1])
			stackRes = append(stackRes[:len(stackRes)-2], strconv.Itoa(first-second))
		} else if val == "*" {
			first, _ := strconv.Atoi(stackRes[len(stackRes)-2])
			second, _ := strconv.Atoi(stackRes[len(stackRes)-1])
			stackRes = append(stackRes[:len(stackRes)-2], strconv.Itoa(first*second))
		} else if val == "/" {
			first, _ := strconv.Atoi(stackRes[len(stackRes)-2])
			second, _ := strconv.Atoi(stackRes[len(stackRes)-1])
			stackRes = append(stackRes[:len(stackRes)-2], strconv.Itoa(first/second))
		} else {
			stackRes = append(stackRes, val)
		}
	}
	ans, _ := strconv.Atoi(stackRes[0])
	return ans
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
