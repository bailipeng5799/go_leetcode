package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	level := 0
	//保存行数奇数行从头到尾输入
	//偶数行从尾到头输入
	for len(queue) != 0{
		i := len(queue)
		//提前根据i的大小申请空间
		//i代表每一层的元素在队列中的长度
		tmp := make([]int,i)
		count := 0
		for ; i > 0; i--{
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil{
				queue = append(queue,node.Left)
			}
			if node.Right != nil{
				queue = append(queue,node.Right)
			}
			if level % 2 ==0{
				tmp[count] = node.Val
				count++
			}else{
				tmp[i-1] = node.Val
			}
		}
		level++
		res = append(res,tmp)

	}
	return res
}
