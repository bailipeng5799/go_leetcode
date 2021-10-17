package main
//利用length来进行平均值的添加
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	list := []*TreeNode{}
	if root == nil {
		return res
	}
	list = append(list,root)
	for len(list) > 0 {
		sum := 0
		length := len(list)
		for i := 0; i < length; i++ {
			node := list[0]
			list = list[1:]
			if node.Left != nil {
				list = append(list,node.Left)
			}
			if node.Right != nil {
				list = append(list,node.Right)
			}
			sum += node.Val
		}
		res = append(res,float64(sum)/float64(length))
	}
	return res

}