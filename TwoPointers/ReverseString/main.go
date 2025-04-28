package main

import (
	"fmt"
)

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {
	fmt.Println([]byte{'h', 'e', 'l', 'l', 'o'})
}
