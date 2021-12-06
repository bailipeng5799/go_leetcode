package main
type LRUCache struct {
	Key []int
	Value []int
	Size int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Key : make([]int,0),
		Value : make([]int,0),
		Size : capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	getValue := -1
	for i := 0; i < len(this.Key); i++ {
		if this.Key[i] == key {
			getValue = this.Value[i]
			this.Key = append(this.Key[:i],append(this.Key[i+1:],this.Key[i])...)
			this.Value = append(this.Value[:i],append(this.Value[i+1:],this.Value[i])...)
			break
		}
	}
	return getValue
}

func (this *LRUCache) Put(key int, value int)  {
	flag := false //标志关键字是否存在
	index := -1
	for i := 0; i < len(this.Key); i++ {
		if key == this.Key[i] {
			flag = true
			index = i
			break
		}
	}
	if len(this.Key) < this.Size { //如果当前长度小于总长度
		if flag { //如果存在那么直接覆盖
			this.Value[index] = value
			this.Key = append(this.Key[:index],append(this.Key[index+1:],this.Key[index])...)
			this.Value = append(this.Value[:index],append(this.Value[index+1:],this.Value[index])...)
		}else {
			this.Key = append(this.Key,key)
			this.Value = append(this.Value,value)
		}
	}else if len(this.Key) == this.Size { //如果长度等于size
		if flag {
			this.Value[index] = value
			this.Key = append(this.Key[:index],append(this.Key[index+1:],this.Key[index])...)
			this.Value = append(this.Value[:index],append(this.Value[index+1:],this.Value[index])...)
		}else {
			this.Key = this.Key[1:]
			this.Value = this.Value[1:]
			this.Key = append(this.Key,key)
			this.Value = append(this.Value,value)
		}

	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
