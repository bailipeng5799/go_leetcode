package main
// 合并两个有序的数组
func merge( A []int ,  m int, B []int, n int ) {
	// write code here
	indexa, indexb := 0, 0
	tmp := 0
	for indexa < m+n {
		if indexa >= m {
			A[indexa] = B[tmp]
			tmp++
		}
		for indexb < n && A[indexa] >= B[indexb] {
			A[indexa], B[indexb] = B[indexb], A[indexa]
			indexb++
		}
		indexa++
		indexb = tmp
	}
}