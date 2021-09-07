package main

func intersection(nums1 []int, nums2 []int) []int {
	m0 := map[int]bool{}
	for _, v := range nums1{
		m0[v] = true
	}
	k := 0
	for _, v := range nums2{
		if m0[v] {
			m0[v] = false
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}