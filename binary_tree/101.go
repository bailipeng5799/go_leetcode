package main
//递归
func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left,root.Right)
}
func dfs(left *TreeNode,right *TreeNode) bool{
	//先判断两个都不为nil
	if left == nil && right == nil {
		return true
	}
	//入果有一个为nil那么返回false
	if left == nil || right == nil {
		return false
	}
	//如果两个值不相同返回false
	if left.Val != right.Val {
		return false
	}
	//因为需要判读是否为镜像所以需要判断左子树的左节点和右子树的右节点，左子树的右节点和右子树的左节点 是否相等
	return dfs(left.Left,right.Right) && dfs(left.Right,right.Left)
}
//迭代
func isSymmetric2(root *TreeNode) bool {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue,root.Left,root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue //这里注意不是直接返回 是进入下一次循环
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue,left.Left,right.Right,right.Left,left.Right)
	}
	return true
}
