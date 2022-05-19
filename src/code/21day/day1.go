package main

import (
	"fmt"
	"sort"
)

/**
双指针
*/

func main() {
	fmt.Println(0x41)
	//nums := []int{0, 2, 2, 0, 3, 3, 4, 4}
	//fmt.Println(removeDuplicates(nums))
	//fmt.Println(removeElement(nums, 2))
	//moveZeroes(nums)
	//fmt.Println(nums)

	//s := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val: 3,
	//				Next: &ListNode{
	//					Val: 3,
	//					Next: &ListNode{
	//						Val: 5,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	//fmt.Println(strStr("mississippi", "issip"))
	//fmt.Println(fourSum([]int{1,0,-1,0,-2,2},0))
}

/**
	18、四数之和
	给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d< n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。



示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSums(nums,4,0,target)
}

//设计一个能求n数之和的API
func nSums(nums []int,n int, start int,target int) [][]int{
	ans:=[][]int{}
	if n<2||len(nums)<n{
		return ans
	}
	if n==2{ //如果n==2的话就直接双指针的两数之和
		left,right:=start,len(nums)-1
		for left<right{
			if nums[left]+nums[right]==target{
				ans=append(ans,[]int{nums[left],nums[right]})
				for left<right&&nums[left]==nums[left+1]{
					left++
				}
				for left<right&&nums[right]==nums[right-1]{
					right--
				}
				left++
				right--
			}else if nums[left]+nums[right]<target{ //
				left++
			}else{
				right--
			}
		}
	}else { //大于两数之和，递归求解
		for i:=start;i<len(nums);i++{
			subRes:=nSums(nums,n-1,i+1,target-nums[i])

			//每一个都添加当前元素
			for j:=range subRes{
				subRes[j]=append(subRes[j],nums[i])
				ans=append(ans,subRes[j])
			}
			//去重
			for i<len(nums)-1&&nums[i+1]==nums[i]{
				i++
			}
		}
	}
	return ans
}

/**
28、实现strStr
给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
输入：haystack = "hello", needle = "ll"
输出：2
输入：haystack = "aaaaa", needle = "bba"
输出：-1
*/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	var slow, fast, count int
	for fast < len(haystack) {
		if haystack[fast] == needle[count] {
			count++
		} else {
			count = 0
			fast = slow
			slow++
		}
		if count == len(needle) {
			return slow
		}
		fast++
	}
	return -1
}

func removeDuplicates(nums []int) int {
	fast := 1
	slow := 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func removeElement(nums []int, val int) int {
	var fast, slow int
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

/**
[1,0,4,1,0,2,2]
*/

func moveZeroes(nums []int) {
	var fast, slow int
	for fast < len(nums) {
		if nums[fast] != 0 {
			if fast != slow {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
		fast++
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	for slow != nil && slow.Next != nil {
		if slow.Val == slow.Next.Val {
			slow.Next = slow.Next.Next
		} else {
			slow = slow.Next
		}
	}
	return head
}
