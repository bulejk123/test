package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBNCA","ABC"))
}

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	//先将要找的字符串 字符和个数存起来
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		//只有在子字符串中的字符才存进去
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		//如果符合条件了
		for check() && l <= r {
			//如果找的比之前找的长则不替换
			if r - l + 1 < len {
				len = r - l + 1
				ansL, ansR = l, l + len
			}
			//从left 开始+1 缩小窗口  直到不符合条件
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}