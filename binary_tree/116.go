package main


type Node struct {
   Val int
   Left *Node
   Right *Node
   Next *Node
}


func connect(root *Node) *Node {
	var res *Node
	if root == nil {
		return res
	}
	list := []*Node{}
	list = append(list,root)
	for len(list) > 0 {
		length := len(list)
		for i := 0; i < length; i++{
			node := list[0]
			list = list[1:]
			if node.Left != nil {
				list = append(list,node.Left)
			}
			if node.Right != nil {
				list = append(list,node.Right)
			}
			if i+1 < length {//当i+1<length是为了避免最后一个元素让其指向nil
				node.Next = list[0]
			}
		}
	}
	return root
}
