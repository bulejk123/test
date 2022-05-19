package main

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func main() {

}

//删除二叉搜索数的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val == key {
		if root.Left == nil{
			return root.Right
		}
		if root.Right == nil{
			return root.Left
		}
		minNode:=getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right,minNode.Val)

	}else if root.Val > key{
		root.Left = deleteNode(root.Left,key)
	}else if root.Val < key {
		root.Right = deleteNode(root.Right,key)
	}
	return root
}
func getMin(node *TreeNode) *TreeNode{
	for node.Left != nil{
		node = node.Left
	}
	return node
}

//二叉树展开
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left,right:=root.Left,root.Right
	root.Left = nil
	root.Right = left
	p:=root
	for p.Right != nil{
		p = p.Right
	}
	p.Right = right
}

//查找目标节点
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if  root.Val == val {
			return root
		}
		if root.Val > val {
			root = root.Left
		}else{
			root = root.Right
		}
	}
	return nil
}

//判断是否是二叉搜索数
func isValidBST(root *TreeNode) bool {
	return isValid(root,nil,nil)
}

func isValid(root *TreeNode,min *TreeNode,max *TreeNode) bool {
	if root == nil{
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValid(root.Left,min,root) && isValid(root.Right,root,max)
}