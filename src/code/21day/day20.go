package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
	var r []int
	for i := left; i <= right; i++ {
		temp := i
		flag := true
		for temp > 0 {
			j := temp % 10
			if j == 0 || i%j != 0 {
				flag = false
				break
			}
			temp=temp/10
		}
		if flag {
			r = append(r,i)
		}
	}
	return r
}

func constructArr(a []int) []int {
	dp := make([]int, len(a))
	sum := 1
	for i := len(a) - 1; i >= 0; i-- {
		sum *= a[i]
		dp[i] = sum
	}
	return dp
}

func maxSubArray(nums []int) int {
	sum, res := -math.MaxInt32, -math.MaxInt32
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	count := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > count {
			count = dp[i]
		}
	}
	return count
}
