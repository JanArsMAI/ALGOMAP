package main

func romanToInt(s string) int {
	d := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	n := len(s)
	for i := 0; i < n; i++ {
		val := d[s[i]]
		if i+1 < n && d[s[i+1]] > val {
			res -= val
		} else {
			res += val
		}
	}
	return res
}

func main() {
}
