package main

import (
	"fmt"
)

func maxNumberOfBalloons(text string) int {
	resMap := make(map[rune]int)
	cnt := 0
	for _, r := range text {
		resMap[r]++
	}
	for {
		bcount, bexist := resMap['b']
		acount, aexist := resMap['a']
		lcount, lexist := resMap['l']
		ocount, oexist := resMap['o']
		ncount, nexist := resMap['n']
		if !bexist || !aexist || !lexist || !oexist || !nexist {
			return 0
		}
		if bcount >= 1 && acount >= 1 && lcount >= 2 && ocount >= 2 && ncount >= 1 {
			cnt++
			resMap['b']--
			resMap['a']--
			resMap['l']--
			resMap['l']--
			resMap['o']--
			resMap['o']--
			resMap['n']--
		} else {
			return cnt
		}
	}
}

func main() {
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}
