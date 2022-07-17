package main

import (
	"strings"
)

func main() {
	str := "alpha alpha alpha alpha alpha alpha alpha alpha"

	println(replaceNth(str, "alpha", "beta", 3))
}

func replaceNth(s, old, new string, n int) string {
	i := 0
	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		//found it
		i += x
		if j == n {
			s = s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
