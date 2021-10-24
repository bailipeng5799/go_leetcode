package main
//669. 修剪二叉搜索树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil { //如果碰到nil直接返回nil节点
		return nil
	}
	if root.Val < low { //如果小于进入右子树
		right := trimBST(root.Right,low,high)
		return right
	}
	if root.Val > high {//如果大于进入左子树
		left := trimBST(root.Left,low,high)
		return left
	}
	root.Left = trimBST(root.Left,low,high) //说明当前在范围内，将左右子树分别递归得到左右节点
	root.Right = trimBST(root.Right,low,high)
	return root             //返回
}
