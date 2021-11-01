package main
func levelOrder( root *TreeNode ) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{}
	stack = append(stack,root)
	res := [][]int{}
	for len(stack) != 0 {
		tmp := []int{}
		length := len(stack)
		for i := 0; i < length; i++{
			node := stack[0]
			stack = stack[1:]
			tmp = append(tmp,node.Val)
			if node.Left != nil {
				stack = append(stack,node.Left)
			}
			if node.Right != nil {
				stack = append(stack,node.Right)
			}
		}
		res = append(res,tmp)
	}
	return res
}
//递归
func levelOrder2( root *TreeNode ) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode,level int)
	dfs = func(root *TreeNode,level int){
		if root == nil {
			return
		}
		if len(res) == level{
			res = append(res,[]int{})
		}
		res[level] = append(res[level],root.Val)
		dfs(root.Left,level+1)
		dfs(root.Right,level+1)
	}
	dfs(root,0)
	return res

}

