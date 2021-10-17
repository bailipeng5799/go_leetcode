package main
//和层序遍历一样只是多了翻转的过程
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	list := []*TreeNode{}
	if root == nil {
		return res
	}
	list = append(list,root)
	for len(list) > 0 {
		tmp := []int{}
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
			tmp = append(tmp,node.Val)
		}
		res = append(res,tmp)
	}
	for i := 0; i < len(res)/2; i++{
		res[i],res[len(res)-1-i]=res[len(res)-1-i],res[i]
	}
	return res
}
