package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var path []int
	recur(root,target,path,&res)
	return res
}
func recur(root *TreeNode,target int,path []int,res *[][]int){
	if root == nil{
		return
	}

	path = append(path,root.Val)
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil{
		//不能直接将path添加到res中因为如果最后一个root.Right如果不为空
		//会修改你最后一次添加的path的最后一个值
		//理解就是当你append一个新元素后切片指向他的底层数组，如果你对这个数组进行修改，这个切片在下一次append
		//一个新的数组之前，值也会被修改
		*res = append(*res, append([]int{}, path...))
	}
	recur(root.Left,target,path,res)
	recur(root.Right,target,path,res)
	path = path[:len(path)-1]
}

