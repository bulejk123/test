package main

import (
	"fmt"
	"sort"
)

func main() {
	//[2,4,0,0,8,1]
	fmt.Println(canReorderDoubled([]int{2,4,0,0,8,1}))
}

func canReorderDoubled1(arr []int) bool {
	m:=make(map[int]int,len(arr))
	for _, num := range arr {
		m[num]++
	}
	//如果值为0的只有奇数个 必然为false
	if m[0]%2==1 {
		return false
	}
	vals := make([]int, 0, len(m))
	for x := range m {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })
	for _, val := range vals {
		if m[2*val] < m[val] {
			return false
		}
		m[2*val]-=m[val]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


/**
0:arr[1] = 2*arr[0]
1:arr[3] = 2*arr[2]
2:arr[5] = 2*arr[4]
 */
func canReorderDoubled(arr []int) bool {
	flag:=false
	for i := 0; i < len(arr) ; i++ {
		for j := 0; j < len(arr) - i -1 ; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1],arr[j] = arr[j],arr[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	var index []int
	for i := 0; i < len(arr)-1; i++ {
		if checkNum(index,i) {
			continue
		}
		flag :=false
		for j := i+1; j < len(arr); j++ {
			if checkNum(index,j) {
				continue
			}
			if arr[i] == 2*arr[j] || 2*arr[i] == arr[j]{
				index = append(index,i,j)
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
func checkNum(a []int,num int)bool{
	for _, n := range a {
		if n == num {
			return true
		}
	}
	return false
}