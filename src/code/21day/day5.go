package main

import "fmt"

func main() {
	fmt.Println(search1([]int{-1,0,3,5,9,12},5))
}

func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	left:=0
	right:=len(nums)-1
	for left <= right{
		mid:=left+(right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid -1
		}else {
			left = mid + 1
		}
	}

	return -1
}
