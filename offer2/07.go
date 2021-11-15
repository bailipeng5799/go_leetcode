package main
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	for i := 0; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val : preorder[0],
				Left : buildTree(preorder[1:len(inorder[:i])+1],inorder[:i]),
				Right : buildTree(preorder[len(inorder[:i])+1:],inorder[i+1:]),
			}
		}

	}
	return nil

}
