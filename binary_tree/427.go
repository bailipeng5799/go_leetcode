package main
type Node struct {
	Val      int
	Children []*Node
}
//将左右子树换为一个循环遍历N叉树
func levelOrder427(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	list := []*Node{}
	list = append(list,root)
	for len(list) > 0 {
		length := len(list)
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := list[0]
			list = list[1:]
			for j := 0; j < len(node.Children); j++ {
				list = append(list,node.Children[j])
			}
			tmp = append(tmp,node.Val)
		}
		res = append(res,tmp)
	}
	return res
}
