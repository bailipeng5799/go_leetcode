package main
func LRU2( operators [][]int ,  k int ) []int {
	res := make([]int,0) //结果切片
	key := make([]int,0) //存放key的切片
	value := make([]int,0)//存放value的切片
	for i := 0; i < len(operators); i++{
		if operators[i][0] == 1{ //表示加入keyvalue中
			if len(key) <= k - 1 { //如果当前长度小于k的长度
				flag := false //用来表示即将要set的数据是否已经存在切片中
				for j := 0;j < len(key); j++{
					if key[j] == operators[i][1]{ //如果已经存在，那么将set的value覆盖并且更新在切片中的顺序
						value[j] = operators[i][2]
						key = append(key[:j],append(key[j+1:],key[j])...)
						value = append(value[:j],append(value[j+1:],value[j])...)
						flag = true //flag为true表示存在并且操作结束
					}
				}
				if !flag{ //如果之前不存在直接添加
					key = append(key,operators[i][1])
					value = append(value,operators[i][2])
				}
			}else{ //如果当前长度等于k 将最久没有使用的元素也就是索引为0的元素删除，并将要添加的元素进行添加
				key = key[1:]
				value = value[1:]
				key = append(key,operators[i][1])
				value = append(value,operators[i][2])
			}
		}
		if operators[i][0] == 2 { //如果查找
			index := -1        //没查找到返回-1
			for j := 0; j < len(key); j++ {
				if operators[i][1] == key[j] { //查找到了将值暂时存储更行keyvalue中的顺序
					index = value[j]
					key = append(key[:j],append(key[j+1:],key[j])...)
					value = append(value[:j],append(value[j+1:],value[j])...)
					break
				}
			}
			res = append(res,index) //将index加入结果集
		}
	}
	return res //返回res
}
