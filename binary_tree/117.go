package main
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect2(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, prev *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur  //每一层的第一个开始节点
			}
			if prev != nil {
				prev.Next = cur //指向右边的节点
			}
			prev = cur //保存前一个节点，用来指向下一个元素
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart //start初始化
	}
	return root
}


