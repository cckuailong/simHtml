package main

import "fmt"

func longestCommonSubsequence(text1, text2 []string) int {
	m, n := len(text1), len(text2)
	up := make([]int, n+2)
	var a, b, c, tmp, maximum int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				tmp = a + 1
			} else {
				tmp = max(b, c)
			}
			if tmp > maximum {
				maximum = tmp
			}
			c = tmp
			a = b
			up[j] = tmp
			b = up[j+1]
		}
		a = 0
		b = up[1]
		c = 0
	}
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	res := longestCommonSubsequence([]string{"aaa","bbb","ccc","dddd"}, []string{"bbb","dddd"})
	fmt.Println(res)
}
