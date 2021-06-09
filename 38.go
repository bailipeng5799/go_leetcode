package main

func permutation(s string) []string {
	res := []string{}//结果的切片
	bytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int){
		if x == len(bytes) - 1{
			res = append(res,string(bytes))
			return
		}
		tmp := map[byte]bool{}
		for i := x;i < len(bytes);i++{
			//如果这个字符没有被使用过则可以用来继续使用如果被使用过根据map的规则为true直接跳过
			if !tmp[bytes[i]]{
				bytes[x],bytes[i] = bytes[i],bytes[x]
				tmp[bytes[x]] = true//把这个固定的字符的bool设置为true当下一次访问时直接跳过可能存在两个相同的字符
									// 其实就是设置第几个字符为true如果存在两个b 那么第二个b在这个位置将不能进行交换
				dfs(x+1)//标记为true后直接进入
				bytes[x],bytes[i] = bytes[i],bytes[x]//如果使用完毕就将两个字符交换回去
			}
		}
	}
	dfs(0)
	return res
}

