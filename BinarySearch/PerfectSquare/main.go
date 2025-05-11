package main

func isPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}
	x := num
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}
