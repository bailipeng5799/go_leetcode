package main
//二叉树的层序遍历

// Definition for a binary tree node.
  type TreeNode struct {
  	Val int
 	Left *TreeNode
    Right *TreeNode
  }

/*
   主要思想就是外层循环控制层数，内存循环控制每一层的几个元素
*/
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil{
		return nil
	}
	qunue := []*TreeNode{}
	qunue = append(qunue,root)
	for len(qunue) > 0{//
		current := []int{}
		tmp := len(qunue)//用来存放这一层有多少个元素个数
		for tmp > 0{
			node := qunue[0]
			qunue = qunue[1:]
			current = append(current,node.Val)
			if node.Left != nil {
				qunue = append(qunue,node.Left)
			}
			if node.Right != nil{
				qunue = append(qunue,node.Right)
			}
			//将队列中本层的所有元素的左子树和右子树同时添加到此队列中
			//由于tmp存储了上一层的元素个数每次将tmp--
			tmp--
		}
		result = append(result,current)//将current添加到结果切片中
	}
	return result
}
//先进行判断是否为空的操作
func DFS(root *TreeNode,level int,res [][]int)[][]int{
	if root == nil{
		return res
	}
	if len(res) == level{//为了判断一共有多少层，给每一层赋值
		res = append(res,[]int{root.Val})
	}else{//添加到所属与的level这个层的切片中
		res[level] = append(res[level],root.Val)
	}
	res = DFS(root.Left,level+1,res)
	res = DFS(root.Right,level+1,res)
	return res
}