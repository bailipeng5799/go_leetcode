package main
//队列
func countNodes(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue,root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue,node.Left)
		}
		if node.Right != nil {
			queue = append(queue,node.Right)
		}
		res++
	}
	return res
}
//digui
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}