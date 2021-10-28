package main
func MySort( arr []int ) []int {
	quickSort(arr,0,len(arr)-1)
	return arr
}
type Stack struct {
	S []*Index
}
type Index struct{
	Left int
	Right int
}
//入栈
func(s *Stack) Push(l,r int) {
	index := Index{
		Left: l,
		Right: r,
	}
	s.S = append(s.S,&index)
}
//出栈
func(s *Stack) Pull() *Index {
	if len(s.S) == 0 {
		return nil
	}
	res := s.S[len(s.S)-1]
	s.S = s.S[:len(s.S)-1]
	return res
}
//非递归
func quickSort2(arr []int) {
	length := len(arr)
	var stack Stack
	stack.Push(0,length-1)//栈中存在的元素其实和递归相似 left 和 right存储的是每次需要比较的界限
	for {
		s := stack.Pull() //先出栈
		if s == nil { //如果栈为空就break
			break
		}
		//和递归思想一样
		if s.Left < s.Right {
			i := s.Left
			j := s.Right
			mid := arr[(s.Left+s.Right)/2]
			for i < j {
				for arr[i] < mid{
					i++
				}
				for arr[j] > mid {
					j--
				}
				if i <= j {
					arr[i],arr[j] = arr[j],arr[i]
					i++
					j--
				}
			}
			if s.Left < j {
				stack.Push(s.Left,j)
			}
			if s.Right > j {
				stack.Push(i,s.Right)
			}
		}
	}
}
func quickSort(arr []int,start,end int) {
	if start < end {
		i := start
		j := end
		mid := arr[(start+end)/2]
		for i < j {
			for arr[i] < mid {
				i++
			}
			for arr[j] > mid {
				j--
			}
			if i <= j {
				arr[i],arr[j] = arr[j],arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(arr,start,j)
		}
		if end > i {
			quickSort(arr,i,end)
		}

	}
}

