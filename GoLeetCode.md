# GoLeetCode

[TOC]





## 问题集

### 1. Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

**Example:**

> Given nums = [2, 7, 11, 15], target = 9,

> Because nums[0] + nums[1] = 2 + 7 = 9,

> return [0, 1].

**代码**

[Two Sum](https://github.com/artong0416/GoLeetCode/blob/master/twoSum.go "twoSum.go")

```go
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
```

### 2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

> Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)

> Output: 7 -> 0 -> 8

**代码**

[Add Two Numbers](https://github.com/artong0416/GoLeetCode/blob/master/addTwoNumbers.go "Add Two Numbers")

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	retv := ret
	carry := 0
	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		tmp := val1 + val2 + carry

		carry = tmp / 10

		ret.Next = &ListNode{tmp % 10, nil}
		ret = ret.Next
	}
	if carry == 1 {
		ret.Next = &ListNode{1, nil}
	} else {
		ret = nil
	}
	return retv.Next
}
```
### 3. Longest Substring Without Repeating Characters

Given a string, find the length of the **longest substring** without repeating characters.

**Examples:**

Given `"abcabcbb"`, the answer is `"abc"`, which the length is 3.

Given `"bbbbb"`, the answer is `"b"`, with the length of 1.

Given `"pwwkew"`, the answer is `"wke"`, with the length of 3. Note that the answer must be a **substring**, `"pwke"` is a *subsequence* and not a substring.

**思路**

使用选定的数据结构记录某字符最近一次出现的位置（可以是数组可以是map），index记录当前子串的起始位置，那么当前无重复子串长度为当前索引减去起始位置。如果某字符不是第一次出现，则将子串其实位置设置为该字符后一位。一次遍历。

**代码**

[Longest Substring Without Repeating Characters](https://github.com/artong0416/GoLeetCode/blob/master/lengthOfLongestSubstring.go"Longest Substring Without Repeating Characters")

```go
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
```