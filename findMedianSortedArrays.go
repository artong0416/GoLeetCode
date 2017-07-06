/*
寻找中位数
*/

package main

import "fmt"

func main()  {
	num1 := []int {1,2,3,5,8}
	num2 := []int {2,6,8,9}
	fmt.Println(findMedianSortedArrays(num1, num2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	suml := l1 + l2
	if suml % 2 == 1 {
		return float64(findKthNum(nums1, nums2, l1, l2, suml/2 + 1))
	} else {
		return float64(findKthNum(nums1, nums2, l1, l2, suml/2) + findKthNum(nums1, nums2, l1, l2, suml/2 + 1)) / 2.0
	}
}

//寻找第k大
func findKthNum(nums1 []int, nums2 []int, m int, n int, k int) int {
	//简化操作位第一个数组比第二个数组短的情况
	if m > n {
		return findKthNum(nums2, nums1, n, m, k)
	}
	//终止条件1，第一个数组为空
	if m == 0 {
		return nums2[k-1]
	}
	//终止条件2，问题变为找最小的
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	//二分查找
	ia := min(k/2, m)
	ib := k - ia
	//当数组1的k/2小于数组2的k/2元素去掉数组1的左边
	if nums1[ia - 1] < nums2[ib - 1] {
		return findKthNum(nums1[ia:m], nums2, m - ia, n, k - ia)
	} else if nums1[ia - 1] > nums2[ib - 1] {
		//当数组1的k/2大于数组2的k/2元素去掉数组2的左边
		return findKthNum(nums1, nums2[ib:n], m, n-ib, k - ib)
	} else {
		//相等，则都是第k
		return nums1[ia - 1]
	}
}

// golang min int
func min(first int, args... int) int {
	for _ , v := range args{
		if first > v {
			first = v
		}
	}
	return first
}