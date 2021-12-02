package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//递归三种
func Preorder(root *TreeNode)[]int {
	res := []int{}
	var Traversal func(node *TreeNode)
	Traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res,node.Val)
		Traversal(node.Left)
		Traversal(node.Right)
	}
	Traversal(root)
	return res
}
func Inorder(root *TreeNode)[]int{
	res := []int{}
	var Traversal func(node *TreeNode)
	Traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		Traversal(node.Left)
		res = append(res,node.Val)
		Traversal(node.Right)
	}
	Traversal(root)
	return res
}
func Postorder(root *TreeNode)[]int{
	res := []int{}
	var Traversal func(node *TreeNode)
	Traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		Traversal(node.Left)
		Traversal(node.Right)
		res = append(res,node.Val)
	}
	return res
}
func Sequence(root *TreeNode)[][]int{
	res := [][]int{}
	return DFS(root,0,res)
}
func DFS(node *TreeNode,level int,res [][]int)[][]int  {
	if node == nil {
		return res
	}
	if len(res) == level {
		res = append(res,[]int{node.Val})
	}else {
		res[level] = append(res[level],node.Val)
	}
	res = DFS(node.Left,level+1,res)
	res = DFS(node.Right,level+1,res)
	return  res
}


//迭代
func preorder2(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack,node)
			res = append(res,node.Val)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
func inorder(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack,node.Left)
			node = node.Left
		}
		node = stack[len(stack)-1]
		res = append(res,node.Val)
		stack = stack[:len(stack)-1]
		node=node.Right
	}
	return res
}
func postorder(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	node := root
	var prev *TreeNode //存放前一个节点
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack,node.Left)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev { //判断右节点是否已经遍历过了
			res = append(res,node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack,node)
			node = node.Right
		}
	}
}
func sequence2(root *TreeNode) [][]int{
	list := []*TreeNode{}
	res := [][]int{}
	if root == nil {
		return res
	}
	list = append(list,root)
	for len(list) > 0 {
		length := len(list)
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := list[0]
			list = list[1:]
			if node.Left != nil{
				list = append(list,node.Left)
			}
			if node.Right != nil {
				list = append(list,node.Right)
			}
			tmp = append(tmp,node.Val)
		}
		res = append(res,tmp)
	}
	return res


}
func main(){

}
