package main
//235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val{ //当前节点的值大于pq的值说明在左面
		return lowestCommonAncestor(root.Left,p,q)
	}
	if root.Val < p.Val && root.Val < q.Val{   //当前节点的值小于pq的值说明在右面
		return lowestCommonAncestor(root.Right,p,q)
	}else{                                   //当前节点的值在pq中间，为最深的祖先
		return root
	}
}
