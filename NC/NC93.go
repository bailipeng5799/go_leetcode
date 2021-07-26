package main
//设计LRU缓存结构
//设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
//set(key, value)：将记录(key, value)插入该结构
//get(key)：返回key对应的value值
//[要求]
//set和get方法的时间复杂度为O(1)
//某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
//当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
//若opt=1，接下来两个整数x, y，表示set(x, y)
//若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
//对于每个操作2，输出一个答案
func LRU( operators [][]int ,  k int ) []int {
	res := make([]int,0,len(operators))
	key := make([]int,k)
	value := make([]int,k)
	for _,v := range operators{
		if v[0] == 1{
			//插入
			//切片末尾表示最新元素
			//先判断个数是否等于缓存大小如果等于删除最旧的元素
			if len(key) == k{
				key = key[1:]
				value = value[1:]
			}
			key = append(key,v[1])
			value = append(value,v[2])
		}else if v[0] == 2{
			//获取
			index := -1
			for i := 0; i < len(key); i++{
				if v[1] == key[i]{
					index = i
					break
				}
			}
			//如果index == -1 说明没有查找到将-1放入结果切片
			if index == -1{
				res = append(res,-1)
			}else{
				//将GET所获取的value添加到结果切片中
				res = append(res,value[index])
				if index < k - 1{
					//如果所要查找的下表index为缓存切片中的元素
					//将其置换在切片末尾
					key = append(key[:index],append(key[index+1:],key[index])...)
					value = append(value[:index],append(value[index+1:],value[index])...)
				}
			}

		}
	}

	return res

}
