package main

import (
	"fmt"
	"math"
)

func main() {
	//[312884470]
	//312884469
	//[873375536,395271806,617254718,970525912,634754347,824202576,694181619,20191396,886462834,442389139,572655464,438946009,791566709,776244944,694340852,419438893,784015530,588954527,282060288,269101141,499386849,846936808,92389214,385055341,56742915,803341674,837907634,728867715,20958651,167651719,345626668,701905050,932332403,572486583,603363649,967330688,484233747,859566856,446838995,375409782,220949961,72860128,998899684,615754807,383344277,36322529,154308670,335291837,927055440,28020467,558059248,999492426,991026255,30205761,884639109,61689648,742973721,395173120,38459914,705636911,30019578,968014413,126489328,738983100,793184186,871576545,768870427,955396670,328003949,786890382,450361695,994581348,158169007,309034664,388541713,142633427,390169457,161995664,906356894,379954831,448138536]
	//943223529


	//[1,1,1,999999999]
	//10
	//fmt.Println(minEatingSpeed([]int{312884470},312884469))
	fmt.Println(shipWithinDays([]int{1,2,3,1,1},4))
}

/**
传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。

 

示例 1：

输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。
示例 2：

输入：weights = [3,2,2,4,1,4], days = 3
输出：6
解释：
船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
第 1 天：3, 2
第 2 天：2, 4
第 3 天：1, 4
示例 3：

输入：weights = [1,2,3,1,1], days = 4
输出：3
解释：
第 1 天：1
第 2 天：2
第 3 天：3
第 4 天：1, 1
 

提示：

1 <= days <= weights.length <= 5 * 104
1 <= weights[i] <= 500

 */
func shipWithinDays(weights []int, days int) int {
	var max,sum int
	for i := 0; i < len(weights); i++ {
		sum +=weights[i]
		if max < weights[i] {
			max= weights[i]
		}
	}
	left:=max
	right:=sum
	for left <= right {
		mid :=left+(right - left) /2
		day:=getDays(weights,mid)
		if day > days {
			left = mid +1
		}else {
			right = mid -1
		}
	}
	return left
}

func getDays(weights []int,load int)int{
	var sum int
	day:=1
	for _, weight := range weights {
		if sum + weight > load {
			sum = 0
			day++
		}
		sum+=weight
	}
	return day
}



/**
珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。  

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

 

示例 1：

输入: piles = [3,6,7,11], H = 8
输出: 4
示例 2：

输入: piles = [30,11,23,4,20], H = 5
输出: 30
示例 3：

输入: piles = [30,11,23,4,20], H = 6
输出: 23

1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9
*/


func minEatingSpeed(piles []int, h int) int {
	var max int
	for i := 0; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}
	if len(piles) == h {
		return max
	}
	left:=1
	right:=max
	for left <= right {
		mid := left + (right - left) / 2
		var sum int
		for _, pile := range piles {
			sum += int(math.Ceil(float64(pile) / float64(mid)))
		}

		if sum <= h {
			right = mid - 1
		}else {
			left = mid +1
		}
	}
	return left
}

func minEatingSpeed1(piles []int, h int) int {
	var max int
	for i := 0; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}
	if len(piles) == h {
		return max
	}

	for i := 1; i <= max; i++ {
		var sum int
		for _, pile := range piles {
			sum += int(math.Ceil(float64(pile) / float64(i)))
		}
		if sum <= h {
			return i
		}
	}
	return 1
}
