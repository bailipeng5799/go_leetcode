package main
func findMode(root *TreeNode) []int {
	res := make([]int,0) // 存放结果的结果集
	count := 1  //计算众数的临时值
	max := 1    //最大的众数
	var prev *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left) //其实还是一个中序遍历
		if prev != nil && prev.Val == node.Val { //判断中序遍历的前一个节点值是否和当前节点值相等
			count++                              //如果相等count++
		}else{
			count = 1                            //否则初始化count为1
		}
		if count >= max {                        //如果count>=之前的存在的最大众数的数量
			if count > max && len(res) > 0 {     //1.此枪矿说明count数量大于之前众数的数量所以直接初始化res
				res = []int{node.Val}
			}else {                             //此情况说明count == max 此众数的数量等于之前众数的数量
				res = append(res,node.Val)
			}
			max = count                         //将max替换
		}
		prev = node                             //前置指针
		travel(node.Right)                      //前序遍历的进入右节点
	}
	travel(root)
	return res
}
