package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//fmt.Println(numberOfLines([]int{4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10},"bbbcccdddaaa"))
	fmt.Println(time.Now().Unix())
	fmt.Println(removeKdigits("1432219",3))
	fmt.Println(time.Now().Unix())
}

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}


func removeKdigits1(num string, k int) string {
	res:=[]byte(num)
	for k > 0 {
		count:=0
		for i := 1; i < len(res) && res[i] >= res[i-1]; i++ {
			count = i
		}
		res=append(res[:count],res[count+1:]...)
		count=0
		for count<len(res) && res[count] == '0'{
			count++
		}
		res=res[count:]
		if len(res)==0{
			return "0"
		}
		k--
	}
	return string(res)
}

func numberOfLines(widths []int, s string) []int {
	result := make([]int, 2)
	result[0], result[1] = 1, 0
	count := 0
	for i := 0; i < len(s); i++ {
		count += widths[s[i]-'a']
		if count <= 100 {
			result[1] = count
		} else {
			result[0]++
			result[1] = widths[s[i]-'a']
			count = result[1]
		}
	}
	return result
}

//各位数字都不同。
//来详解一下
//dp[i]=dp[i-1]+(dp[i-1]-dp[i-2])*(10-(i-1));
//加上dp[i-1]没什么可说的，加上之前的数字
//dp[i-1]-dp[i-2]的意思是我们上一次较上上一次多出来的各位不重复的数字。以n=3为例，n=2已经计算了0-99之间不重复的数字了，我们需要判断的是100-999之间不重复的数字，那也就只能用10-99之间的不重复的数去组成三位数，而不能使用0-9之间的不重复的数，因为他们也组成不了3位数。而10-99之间不重复的数等于dp[2]-dp[1]。
//当i=2时，说明之前选取的数字只有
//1位，那么我们只要与这一位不重复即可，所以其实有9(10-1)种情况（比如1，后面可以跟0,2,3,4,5,6,7,8,9）。
//当i=3时，说明之前选取的数字有2位，那么我们需要与2位不重复，所以剩余的
//有8（10-2）种（比如12，后面可以跟0,3,4,5,6,7,8,9）

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 10
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + (dp[i-1]-dp[i-2])*(10-(i-1))
	}
	return dp[n]
	//switch n {
	//case 0:
	//	return 1
	//case 1:
	//	return 10
	//case 2:
	//	return 91
	//case 3:
	//	return 739
	//case 4:
	//	return 5275
	//case 5:
	//	return 32491
	//case 6:
	//	return 168571
	//case 7:
	//	return 712891
	//case 8:
	//	return 2345851
	//case 9:
	//	return 5611771
	//default:
	//	return 8877691
	//}

}
