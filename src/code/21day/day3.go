package main

import "fmt"

/**
	差分数组
 */

func main() {
	trips := [][]int{{2, 2, 6}, {2, 4, 7},{8,6,7}}
	fmt.Println(carPooling1(trips, 11))

	//bookings:=[][]int{{1,2,10}, {2, 3, 20},{2,5,25}}
	//fmt.Println(corpFlightBookings(bookings,5))
}


/**
这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

 

示例 1：

输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]
示例 2：

输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]
解释：
航班编号        1   2
预订记录 1 ：   10  10
预订记录 2 ：       15
总座位数：      10  25
因此，answer = [10,25]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/corporate-flight-bookings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func corpFlightBookings(bookings [][]int, n int) []int {
	preDiff:=make([]int,n)
	for i := 0; i < len(bookings); i++ {
		preDiff[bookings[i][0]-1] += bookings[i][2]
		if bookings[i][1] < n {
			preDiff[bookings[i][1]] -= bookings[i][2]
		}
	}
	res:=make([]int,n)
	res[0]=preDiff[0]
	for i := 1; i < len(preDiff) ; i++ {
		res[i] = res[i-1] + preDiff[i]
	}
	return res
}

/**
这儿有一份乘客行程计划表trips[][]，其中trips[i] = [num_passengers, start_location, end_location]包含了第 i 组乘客的行程信息：

必须接送的乘客数量；
乘客的上车地点；
以及乘客的下车地点。
这些给出的地点位置是从你的初始出发位置向前行驶到这些地点所需的距离（它们一定在你的行驶方向上）。

请你根据给出的行程计划表和车子的座位数，来判断你的车是否可以顺利完成接送所有乘客的任务（当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false）。



示例 1：

输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false
示例 2：

输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true
示例 3：

输入：trips = [[2,1,5],[3,5,7]], capacity = 3
输出：true
示例 4：

输入：trips = [[3,2,7],[3,7,9],[8,3,9]], capacity = 11
输出：true

*/


func carPooling1(trips [][]int, capacity int)bool {
	var lenght int
	for i := 0; i < len(trips); i++ {
		if lenght < trips[i][2] {
			lenght = trips[i][2]
		}
	}
	m:=make([]int,lenght+1)
	for _, trip := range trips {
		m[trip[1]] += trip[0]
		m[trip[2]] -= trip[0]
	}
	var sum int
	for _, v := range m {
		sum += v
		if sum > capacity {
			return false
		}
	}
	return true
}

func carPooling(trips [][]int, capacity int) bool {
	var max int
	for i := 0; i < len(trips); i++ {
		if max < trips[i][2] {
			max = trips[i][2]
		}
	}
	preDiff := make([]int, max+1)
	//preDiff[0] = trips[0][0]
	for i := 0; i < len(trips); i++ {
		preDiff[trips[i][1]] += trips[i][0]
		preDiff[trips[i][2]] -= trips[i][0]
	}

	res := make([]int, len(preDiff))
	res[0] = preDiff[0]
	if res[0] > capacity {
		return false
	}
	for i := 1; i < len(preDiff); i++ {
		res[i] = preDiff[i] + res[i-1]
		if res[i] > capacity {
			return false
		}
	}

	return true
}
