package main
func verifyPostorder(postorder []int) bool {
	return recur(postorder,0,len(postorder)-1)
}
//因为这个数组的最后一位是根节点，搜索树满足右子树大于根节点，左子树小于
func recur(root[]int,left,rootindex int)bool{
	if left >= rootindex{
		return true
	}
	tmp := left
	for root[tmp]<root[rootindex]{
		tmp++
	}
	rightindex := tmp
	for root[tmp] > root[rootindex]{
		tmp++
	}
	return tmp==rootindex && recur(root,left,rightindex-1) && recur(root,rightindex,rootindex-1)
}

