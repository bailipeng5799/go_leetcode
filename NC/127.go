package main

/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS( str1 string ,  str2 string ) string {
	// write code here
	maxlength := 0 //记录最长公共字串的长度
	maxlengthindex := 0 //记录锁存储的最长长度的最后一个元素的下标  可以根据maxlength来得到开始的地方
	dp := make([][]int,len(str1)+1) //给这个二维数组申请空间 多申请了一个
	//因为用dp[i+1][j+1]来存储以ij结尾的字符串长度
	for i := 0; i <= len(str1); i++ {
		dp[i] = make([]int,len(str2)+1)
	}
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j] + 1 //如果等于就+1
				if dp[i+1][j+1] > maxlength { //判断与最大值关系
					maxlength = dp[i+1][j+1]
					maxlengthindex = i
				}
			}else {  //如果不等于就设置为0
				dp[i+1][j+1] = 0
			}
		}
	}

	return str1[maxlengthindex-maxlength+1:maxlengthindex+1]

}
