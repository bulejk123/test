package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(checkS("()[{]}"))
}

func checkS(s string)bool{
	for strings.Contains(s,"{}") || strings.Contains(s,"[]") || strings.Contains(s,"()"){
		s=strings.Replace(s,"()","",-1)
		s=strings.Replace(s,"[]","",-1)
		s=strings.Replace(s,"{}","",-1)
	}
	return s==""
}

func longestCommonPrefix(strs []string) string {
	res := ""
	index := 0
	for {
		if index >= len(strs[0]) {
			break
		}
		s := strs[0][index]
		var count int
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) {
				break
			}
			if strs[i][index] != s {
				break
			}
			count++
		}
		if count!= len(strs)-1 {
			break
		}
		res = res + string(s)
		index++
	}
	return res
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		right--
		left++
	}
	return true
}

func romanToInt(s string) int {
	var num int
	left, right := 0, 1
	for left < len(s) {
		if right >= len(s) {
			right = left
		}
		temp := string(s[left]) + string(s[right])
		curNum := getNum(temp)
		if curNum > 0 {
			num += curNum
			left += 2
			right += 2
			continue
		}
		curNum = getNum(string(s[left]))
		num += curNum
		left++
		right++
	}
	return num
}

func getNum(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	case "IV":
		return 4
	case "IX":
		return 9
	case "XL":
		return 40
	case "XC":
		return 90
	case "CD":
		return 400
	case "CM":
		return 900
	default:
		return 0
	}
}
