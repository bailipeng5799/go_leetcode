package main
type TreeNode struct {
   Val int
  Left *TreeNode
  Right *TreeNode
}

var pre []int
var in []int
var post []int
/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders( root *TreeNode ) [][]int {
	pre,in,post=nil,nil,nil
	preOrder(root)
	inOrder(root)
	postOrder(root)
	return [][]int{pre,in,post}

}
func preOrder(root *TreeNode){
	if root == nil{
		return
	}
	pre = append(pre,root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}
func inOrder(root *TreeNode){
	if root == nil{
		return
	}
	inOrder(root.Left)
	in = append(in,root.Val)
	inOrder(root.Right)
}
func postOrder(root *TreeNode){
	if root == nil{
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	post = append(post,root.Val)
}
