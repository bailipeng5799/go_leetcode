package main
func levelOrder(root *TreeNode) []int {
	//模拟一个队列
	//核心思想就是BFS
	//先入队根结点然后将队列中的第一个节点出队
	//并将其的val值append到结果切片中
	//并且判断他的左右子树是否存在，存在将其添加到树的队列中
	//以此作为循环如果队列为空说明全部出队循环结束

	//核心就是每一次将其节点出队，并将他的左右子树入队可以得到按照层序遍历的结果
	queue := []*TreeNode{}
	res := []int{}
	if root == nil{
		return res
	}
	queue = append(queue,root)
	for len(queue) != 0 {
		tmp := queue[0]
		res = append(res,tmp.Val)
		queue = queue[1:]
		if tmp.Left != nil{
			queue = append(queue,tmp.Left)
		}
		if tmp.Right != nil{
			queue = append(queue,tmp.Right)
		}
	}
	return res
}
