package main
type TreeNode struct {
	  Val int
	  Left *TreeNode
	  Right *TreeNode
}
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil{
		return false
	}
	var result bool
	//如果A的值等于B的值那么根结点值相等进入函数
	if A.Val == B.Val{
		result = compare(A,B)//注意不可以直接返回因为可能会有重复的与B.Val相同的值得到true
	}
	if isSubStructure(A.Left,B) || isSubStructure(A.Right,B){
		return true
	}
	return result
}
//这个比较函数是找到根结点值相等后查询其余节点是否相等
func compare(A, B *TreeNode)bool{
	if B == nil{
		return true
	}
	//A==nil说明A没有此节点，而b有说明不是子结构
	if A == nil{
		return false
	}
	return A.Val == B.Val && compare(A.Left,B.Left) && compare(A.Right,B.Right)
}
