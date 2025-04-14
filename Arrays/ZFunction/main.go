package main

func isSubsequence(s string, t string) bool {
	withSent := s + "$" + t
	n := len(withSent)
	zlist := make([]int, 0, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			zlist[i] = min(r-i+1, zlist[i-l])
		}
		for i+zlist[i] < n && s[zlist[i]] == s[i+zlist[i]] {
			zlist[i] += 1
		}
		if i+zlist[i]-1 > r {
			l, r = i, i+zlist[i]-1
		}
	}
	for _, val := range zlist {
		if val == len(t) {
			return true
		}
	}
	return false
}

func main() {

}
