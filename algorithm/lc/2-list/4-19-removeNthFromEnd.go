package main

import (
	"fmt"

	node "gopl/algorithm/lc/common"
)

/**
description:给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/

func main() {
	n7 := &node.ListNode{Val: 7, Next: nil}
	n6 := &node.ListNode{Val: 6, Next: n7}
	n5 := &node.ListNode{Val: 5, Next: n6}
	n4 := &node.ListNode{Val: 4, Next: n5}
	n3 := &node.ListNode{Val: 3, Next: n4}
	n2 := &node.ListNode{Val: 2, Next: n3}
	n1 := &node.ListNode{Val: 1, Next: n2}
	rel := removeElements1(n1, 2)
	fmt.Println(rel)
}

func removeElements1(head *node.ListNode, val int) *node.ListNode {
	dummyHead := &node.ListNode{Next: head}
	fmt.Println("dummy", dummyHead.Next.Next)
	fmt.Println("head", head.Next)
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
