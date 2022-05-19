package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(rotateString1("abcde", "abced"))
	fmt.Println(nthNumber(11))
}

func nthNumber(n int)int{
	dp:=make([]int,n+1)
	dp[1] = 1
	n1,n2,n3:=1,1,1
	for i:=2;i<=n;i++{
		min:=dp[n1]*2
		if dp[n2]*3 < min {
			min=dp[n2]*3
		}
		if dp[n3]*5 < min {
			min=dp[n3]*5
		}
		dp[i] = min
		if min==dp[n1]*2 {
			n1++
		}
		if min==dp[n2]*3 {
			n2++
		}
	   if min==dp[n3]*5 {
			n3++
		}
	}
	return dp[n]
}

func nthUglyNumber(n int) int {
	num:=1
	for i := 0; i < n; i++ {
		flag:=false
		for isUgly(num) {
			num++
			flag = true
		}
		if flag {
			continue
		}

	}
	return num
}

func isUgly(n int) bool {

	if n == 0 {
		return false
	}
	for n!=1 {
		if n%2==0 {
			n/=2
		}else if n%3==0 {
			n/=3
		}else if n%5==0 {
			n/=5
		}
		if n!=1 && n%2!=0 && n%3!=0&&n%5!=0{
			return false
		}
	}
	return true
}


func rotateString1(s string, goal string) bool {
	 return len(s)==len(goal)&&strings.Contains(s+s,goal)

}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		temp := string(s[i])
		index := i
		count := 1
		for count < len(s) {
			index++
			if index >= len(s) {
				index = 0
			}
			temp += string(s[index])
			count++
		}
		if temp == goal {
			return true
		}
	}
	return false
}
