/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

package main

import "fmt"

func main() {
	l1 := &ListNode{2, nil}
	//l1.Next = &ListNode{4, nil}
	//l1.Next.Next = &ListNode{5, nil}

	l2 := &ListNode{8, nil}
	//l2.Next = &ListNode{6, nil}
	//l2.Next.Next = &ListNode{3, nil}

	fmt.Println(l1, l2)
	res := addTwoNumbers(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

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
