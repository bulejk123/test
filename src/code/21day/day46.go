package main

import "fmt"

func main() {
	fmt.Println(generate(30))
}

func getRow(rowIndex int) []int {
	var arr []int
	for i := 0; i <= rowIndex; i++ {
		if i==0 {
			arr=[]int{1}
		}else if i==1{
			arr=[]int{1,1}
		}else {
			temp:=[]int{1}
			for j := 1; j < i; j++ {
				temp=append(temp,res[i-1][j-1]+res[i-1][j])
			}
			temp=append(temp,1)
			arr = temp
		}
	}

	return arr
}

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		if i==0 {
			temp:=[]int{1}
			res = append(res,temp)
		}else if i==1{
			temp:=[]int{1,1}
			res = append(res,temp)
		}else {
			temp:=[]int{1}
			for j := 1; j < i; j++ {
				temp=append(temp,res[i-1][j-1]+res[i-1][j])
			}
			temp=append(temp,1)
			res = append(res,temp)
		}
	}
	return res
}

