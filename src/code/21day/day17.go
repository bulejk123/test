package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(10))
}

var result [][]string

func solveNQueens(n int) [][]string {
	result = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack1(queens, n, 0, columns, diagonals1, diagonals2)
	return result
}

//回溯
func backtrack1(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		result = append(result, board)
		return
	}
	for i := 0; i < n; i++ {
		//这一列上有皇后 continue
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack1(queens, n, row + 1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

func permute(nums []int) (res [][]int) {
	ln := len(nums)
	flags := make([]bool, len(nums))
	var trace func(i int, tmp []int)
	trace = func(k int, tmp []int) {
		if k == ln {
			part := make([]int, len(tmp))
			copy(part, tmp)
			res = append(res, part)
			return
		}
		for i := 0; i < len(nums); i++ {
			if flags[i] {
				continue
			}
			flags[i] = true
			tmp = append(tmp, nums[i])
			trace(k+1, tmp)
			tmp = tmp[:len(tmp)-1]
			flags[i] = false
		}
		return
	}
	trace(0, []int{})
	return
}

var res [][]int

func permute1(nums []int) [][]int {
	res = [][]int{}
	track:=[]int{}
	backtrack(nums,track)
	return res
}

func backtrack(nums,track []int){
	if len(nums) == len(track) {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res,temp)
	}
	/**
	for 选择 in 选择列表:
	    # 做选择
	    将该选择从选择列表移除
	    路径.add(选择)
	    backtrack(路径, 选择列表)
	    # 撤销选择
	    路径.remove(选择)
	    将该选择再加入选择列表
	 */
	for i := 0; i < len(nums); i++ {
		if isContains(track,nums[i]){
			continue
		}
		track = append(track,nums[i])
		backtrack(nums,track)
		track = track[:len(track)-1]
	}
}

func isContains(track []int,num int)bool{
	for i := 0; i < len(track); i++ {
		if track[i] == num {
			return true
		}
	}
	return false
}

