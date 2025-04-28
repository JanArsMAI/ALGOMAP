package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		leftLower := toLower(s[l])
		rightLower := toLower(s[r])
		if leftLower != rightLower {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func main() {
	fmt.Println(isPalindrome("0P"))
}
