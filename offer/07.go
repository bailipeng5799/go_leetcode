package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
//此题主体思想利用递归，先通过前序遍历拿到root结点，在通过root结点在中序遍历中确定位置
//从而确定左子树和右子树的长度，然后通过递归求解
func buildTree(preorder []int, inorder []int) *TreeNode {
	for i := 0; i<len(preorder); i++ {
		if inorder[i] == preorder[0]{//确定根节点
			return &TreeNode{
				Val: preorder[0],//root节点的值等于根节点的值
				Left: buildTree(preorder[1:len(inorder[:i])+1],inorder[:i]),//左子树先序遍历为preorder[1:len(inorder[:i])+1]，从根节点的中序遍历根节点之前的元素为作字数元素的个数得到
				Right: buildTree(preorder[len(inorder[:i])+1:],inorder[i+1:]),
			}
		}
	}
	return nil
}
