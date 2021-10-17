package main
//层序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//队列
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{} //结果集
	list := []*TreeNode{} //队列利用先进先出来判断
	if root == nil { //root为空直接返回
		return res
	}
	list = append(list,root)
	for len(list) > 0 {
		length := len(list) //用来判断每一层有多少个元素  作为下面循环中使用的条件
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := list[0]   //每次读取第一个元素并且删除
			list = list[1:]
			if node.Left != nil {
				list = append(list,node.Left)
			}
			if node.Right != nil {
				list = append(list,node.Right)
			}
			tmp = append(tmp,node.Val)
		}
		res = append(res,tmp)
	}

	return res
}

//递归
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	res = DFS(root,0,res)
	return res
}
func DFS(root *TreeNode,level int,res [][]int) [][]int {
	if root == nil { //level代表层数
		return res
	}
	if len(res) == level{ //此时有一个元素进入了一个新的层，需要给此层创建空间
		res = append(res,[]int{root.Val})
	}else{ //在已经存在的层中
		res[level] = append(res[level],root.Val)
	}
	res = DFS(root.Left,level+1,res)
	res = DFS(root.Right,level+1,res)
	return res
}





