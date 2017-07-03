/*
功能：最长无重复子串
*/

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}

func lengthOfLongestSubstring(s string) int {
	//保存某字符上一次出现的位置 未出现为 -1
	loc := [256]int{}
	for i := 0; i < 256; i++ {
		loc[i] = -1
	}
	//当前子串起始位置
	index := -1
	//最大长度
	max := 0
	for i := 0; i < len(s); i++ {
		//当前字符出现过，修改子串起始位置到该字符上一次出现的后一个位置（位置从-1开始）
		if loc[s[i]] > index {
			index = loc[s[i]]
		}
		//取最大值
		if i-index > max {
			max = i - index
		}
		//记录字符最近一次出现的位置
		loc[s[i]] = i
	}
	return max
}
