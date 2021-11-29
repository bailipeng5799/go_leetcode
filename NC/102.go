package main
//在二叉树中找到两个节点的最近公共祖先
func lowestCommonAncestor( root *TreeNode ,  o1 int ,  o2 int ) int {
	// write code here
	return dfs(root,o1,o2).Val
}
//递归找到这个元素就返回当left和right都存在的时候就返回当前root代表在当前root的左子树和右子树找到了目标节点
func dfs(root *TreeNode, p, q int) *TreeNode{
	if root == nil || root.Val == p || root.Val == q {
		return root
	}
	left := dfs(root.Left,p,q)
	right := dfs(root.Right,p,q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
