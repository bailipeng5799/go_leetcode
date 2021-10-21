package main
//700. 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	var res *TreeNode
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		res = searchBST(root.Left,val)
	}else if root.Val < val {
		res = searchBST(root.Right,val)
	}
	return res
}
