package main
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {   //如果AB中有任意一个为nil返回false
		return false
	}
	if compare(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B){ //进行递归操作
		return true
	}
	return false
}
func compare(a,b *TreeNode) bool {
	if b == nil { //如果b == nil说明判断完成
		return true
	}
	if a == nil  { //如果a == nil 说明出界因为此时b != nil 值不相等也返回false
		return false
	}

	return a.Val == b.Val && compare(a.Left,b.Left) && compare(a.Right,b.Right)
}
