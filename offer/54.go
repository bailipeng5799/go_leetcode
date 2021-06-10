package main
func kthLargest(root *TreeNode, k int) int {
	//二叉搜索树的中序遍历的倒序是一个递减序列
	//每遍历一次k--一下当遍历到k==0时就是所需要的节点
	res := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil{
			return
		}
		dfs(root.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0{
			res = root.Val
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}

