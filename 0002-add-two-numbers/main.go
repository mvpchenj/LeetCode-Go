package main

import (
	"fmt"
)

/**
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

输入：l1 = [0], l2 = [0]
输出：[0]

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
思路
	难点1，从尾往前算
	难点2，长度不一致
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Printf("%+v \n", generateListNode([]int{2, 4, 3}))
	return
	s1 := []int{9, 9, 9, 9, 9, 9, 9}
	s2 := []int{9, 9, 9, 9}
	l1 := generateListNode(s1)
	l2 := generateListNode(s2)
	fmt.Println(addTwoNumbers(l1, l2))
}

// generateListNode 342 逆序slice就是[2,4,3], 组成listNode的时候也是2，4，3
func generateListNode(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	// 构建最后一个Node
	list := &ListNode{
		Val:  s[len(s)-1],
		Next: nil,
	}
	for i := len(s) - 1; i > 0; i-- {
		subNode := &ListNode{
			Val:  s[i-1],
			Next: list,
		}
		subNode.Next = list
	}
	return list
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{}
}
