package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next,list2)
		return list1
	}else{
		list2.Next = mergeTwoLists(list1,list2.Next)
		return list2
	}
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root ==nil{
		return res
	}
	q:=[]*Node{root}
	for q!=nil{
		level:=[]int{}
		temp:=q
		q=nil
		for _, node := range temp {
			level=append(level,node.Val)
			q= append(q,node.Children...)
		}
		res = append(res,level)
	}
	return res
}
