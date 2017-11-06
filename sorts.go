/*
时间：2017/11/6
功能：golang 各种排序算法
*/

package main

import "fmt"

func main() {
	fmt.Println("排序算法汇总")
}

/*
  #冒泡排序:
  重复比较相邻两个元素，顺序错误就叫唤，重复至不在需要叫唤。
  冒泡就是说最小或者最大的元素回慢慢浮到数列的顶端
*/
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

/*
  #选择排序:
  遍历寻找最小或者最大元素放在最前面，直至遍历结束
*/
func selectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

/*
  #插入排序:
  将第一个元素看做有序队列，剩下的作为待排序列，每次去一个插入到已排序队列的指定位置
*/
func insertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

/*
  #希尔排序:
  是插入排序的一种更高效的改进版本（非稳定）
  先将整个带排序的记录序列分割成若干子序列分别进行直接插入排序。待整个序列基本有序后再对全体直接插入排序。
*/
func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}

/*
  #归并排序
  分治法的典型应用。需要额外空间。
*/
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}
func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

/*
  #快速排序
  快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。
  简单描述 选基准 大的放一边小的放一边
*/
func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}
func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
  #堆排序
  利用数据结构--堆。大顶堆和小顶堆。
*/
func heapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}
func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}
func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
  #计数排序
  计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。
  作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
*/
func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组
	sortedIndex := 0
	length := len(arr)
	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return arr
}

/*
  #桶排序
  桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
*/

/*
  #基数排序
  基数排序是一种非比较型整数排序算法，其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较。
*/
