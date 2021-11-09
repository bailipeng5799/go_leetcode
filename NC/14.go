package main
//按之字形顺序打印二叉树
//和层序遍历相似
func Print( pRoot *TreeNode ) [][]int {
	// write code here
	res := [][]int{}
	level := 0 //如果偶数层就翻转以下
	if pRoot == nil {
		return res
	}
	queue := []*TreeNode{}
	queue  = append(queue,pRoot)
	for len(queue) > 0 {
		tmp := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
			tmp = append(tmp,node.Val)
		}
		if level % 2 == 1{ //如果是奇数层进行切片的翻转
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i],tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1],tmp[i]
			}
		}
		level++
		res = append(res,tmp)
	}
	return res
}

