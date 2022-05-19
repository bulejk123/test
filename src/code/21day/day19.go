package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(fib(7))
	fmt.Println(coinChange([]int{1,2,5},11))
}

func coinChange(coins []int, amount int) int {
	nums:=make([]int,amount+1)
	for i := 0; i < len(nums); i++ {
		nums[i] = -2
	}

	return dp(nums,coins,amount)
}

func dp(nums,coins []int,amount int)int{
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if nums[amount] !=-2 {
		return nums[amount]
	}
	min:=math.MaxInt32
	for _, coin := range coins {
		sub:=dp(nums,coins,amount - coin)
		if sub == -1 {
			continue
		}
		if min > sub +1 {
			min = sub +1
		}
	}
	if min==math.MaxInt32 {
		nums[amount] = -1
	}else {
		nums[amount] = min
	}
	return nums[amount]
}


func fib(n int) int {
	if n == 0{
		return 0
	}
	if n==1 || n==2 {
		return 1
	}
	pre,cur:=1,1
	for i := 3; i <= n ; i++ {
		sum :=pre+cur
		pre = cur
		cur = sum
	}
	return cur
}

func fib2(n int) int {
	nums :=make([]int,n+1)
	return helper(nums,n)
}

//备忘录 将之前算的f（n）存起来 ，不要重复计算
func helper(nums []int,n int)int{
	if n == 0{
		return 0
	}
	if n==1 || n==2 {
		return 1
	}
	if nums[n] != 0 {
		return nums[n]
	}
	nums[n] = helper(nums,n-1)+ helper(nums,n-2)
	return nums[n]
}

func fib1(n int) int {
	if n == 0{
		return 0
	}
	if n==1 || n==2 {
		return 1
	}
	return fib1(n-1)+fib1(n-2)
}
