package main

/**
	前缀和
 */

func main() {
	//nodeList:=deleteDuplicates(s)
	//fmt.Println(nodeList)
	//fmt.Println(subarraySum1([]int{1,-1,0,1,-1},0))
}

type NumMatrix struct {
	preSum [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum[i] = make([]int, len(matrix[i])+1)
		for j := 0; j < len(matrix[i]); j++ {
			sum[i][j+1] = matrix[i][j] + sum[i][j]
		}
	}
	return NumMatrix{
		sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i <= row2; i++ {
		sum += this.preSum[i][col2+1] - this.preSum[i][col1]
	}
	return sum
}

/**
[1,-1,0,1,-1,0]
*/
func subarraySum1(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	var sum, count int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := m[sum-k]; ok {
			count += m[sum-k]
		}
		m[sum] += 1
	}

	return count
}

func subarraySum(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		var jSum int
		if nums[i] == k {
			count++
		}
		for j := i + 1; j < len(nums); j++ {
			jSum += nums[j]
			if nums[i]+jSum == k {
				count++
			}
		}
	}
	return count
}
