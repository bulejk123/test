package main

import (
	"fmt"
)

func main() {

	fmt.Println(shortestToChar("asdbfc",'d'))
}


func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	idx := -n
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		ans[i] = i - idx
	}

	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maximumWealth(accounts [][]int) int {
	max:=0
	for _, account := range accounts {
		tempMax:=0
		for _, val := range account {
			tempMax+=val
		}
		if tempMax > max {
			max = tempMax
		}
	}
	return max
}
