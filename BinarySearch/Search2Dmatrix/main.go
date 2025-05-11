package main

import "fmt"

func BinSearch(lst []int, target int) bool {
	l := 0
	r := len(lst) - 1
	for l <= r {
		mid := l + (r-l)/2
		if lst[mid] == target {
			return true
		} else if lst[mid] > target {
			r = mid - 1
		} else if lst[mid] < target {
			l = mid + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	k := 0
	for i, _ := range matrix {
		if matrix[i][0] <= target {
			k = i
		}
	}
	return BinSearch(matrix[k], target)
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
}
