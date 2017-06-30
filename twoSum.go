/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

//用哈希表，存储每对应的下标，注意注意处理有相同数字的情况

package main

import (
	"fmt"
)

func main()  {

	fmt.Println(twoSum([]int{3,1,2,4}, 6))
}

func twoSum(nums []int, target int) []int {
	//创建Hash存储
	mindex := make(map[int]int, len(nums))
	//初始化
	for index, val := range nums{
		mindex[val] = index
	}

	//查找
	for _, val := range nums{
		if _,ok := mindex[target - val]; ok {
			//可能有相同元素
			if mindex[val] ==  mindex[target - val] {
				for index, val2 := range nums  {
					if val2 == val && index !=  mindex[val] {
						return []int{index, mindex[val]}
					}
				}
				continue
			}
			return []int{mindex[val], mindex[target - val]}
		}
	}
	return nil
}