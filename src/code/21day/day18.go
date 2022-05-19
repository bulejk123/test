package main

import (
	"fmt"
	"time"
)

func main() {
	//[3,3,10,2,6,5,10,6,8,3,2,1,6,10,7,2]
	//6
	//[10,1,10,9,6,1,9,5,9,10,7,8,5,2,10,8]
	//11
	fmt.Println(time.Now())
	//fmt.Println(canPartitionKSubsets([]int{10, 1, 10, 9, 6, 1, 9, 5, 9, 10, 7, 8, 5, 2, 10, 8}, 11))
	fmt.Println(generateParenthesis(3))
	fmt.Println(time.Now())
}

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
*/
var s []string

func generateParenthesis(n int) []string {
	s = []string{}

	backtrack20( "", 0, 0, n)
	return s
}

func backtrack20(str string, l, r, n int) {
	if l > n || r > n || r > l {
		return
	}
	if l == n && r == n {
		s = append(s,str)
	}
	backtrack20(str+"(",l+1,r,n)
	backtrack20(str+")",l,r+1,n)
	return
}

//桶
func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	used := make([]bool, len(nums))
	target := sum / k
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return backtrack19(nums, k, 0, 0, target, used)
}

func backtrack19(nums []int, k, bucket, index, target int, used []bool) bool {
	if k == 0 {
		//桶都装满了
		return true
	}
	if bucket == target {
		//装满了当前桶，开始递归下一个桶  也是从index=0开始选数字
		return backtrack19(nums, k-1, 0, 0, target, used)
	}
	for i := index; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if bucket+nums[i] > target {
			continue
		}
		used[i] = true
		bucket += nums[i]
		if backtrack19(nums, k, bucket, i+1, target, used) {
			return true
		}
		used[i] = false
		bucket -= nums[i]
	}
	return false
}

//一直超时！
func canPartitionKSubsets1(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	bucket := make([]int, k)
	target := sum / k

	flag := false
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return backtrack18(nums, bucket, 0, target)
}

func backtrack18(nums, bucket []int, index, target int) bool {
	if len(nums) == index {
		for _, value := range bucket {
			if value != target {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(bucket); i++ {
		if bucket[i]+nums[index] > target {
			continue
		}
		bucket[i] += nums[index]
		if backtrack18(nums, bucket, index+1, target) {
			return true
		}
		bucket[i] -= nums[index]
	}
	return false
}
