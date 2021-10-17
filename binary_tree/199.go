package main

//length 用来判断最后一个节点是否为右视图所能看到的节点
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	list := []*TreeNode{}
	list = append(list,root)
	for len(list) > 0 {
		length := len(list)
		for i := 0;i < length; i++{
			node := list[0]
			list = list[1:]
			if node.Left != nil{
				list = append(list,node.Left)
			}
			if node.Right != nil{
				list = append(list,node.Right)
			}
			if i == length-1 {
				res = append(res,node.Val)
			}
		}
	}
	return res

}
