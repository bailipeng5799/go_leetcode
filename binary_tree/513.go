package main
//513.找树左下角的值
func findBottomLeftValue(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue,root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			//i == 0 说明是最左边的元素
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
		}
	}
	return res
}
