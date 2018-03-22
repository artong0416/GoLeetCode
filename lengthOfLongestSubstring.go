/*
功能：最长无重复子串
*/

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("as"))
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

//map解法
func lengthOfLongestSubstring1(s string) int {

	indexmap := make(map[uint8]int)
	//当前子串起始位置
	index := -1
	//最大长度
	max := 0
	for i:=0 ; i<len(s); i++ {
		indexmap[s[i]] = -1
	}
	for i:=0 ; i<len(s); i++ {
		//出现过
		if (indexmap[s[i]])  > index {
			index = indexmap[s[i]]
		}
		//取最大
		if i - index  > max {
			max = i - index
		}
		indexmap[s[i]] = i
	}
	return max
}
