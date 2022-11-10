package main

import (
	"fmt"

	node "gopl/algorithm/lc/common"
)

/**
description:给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

// 双指针1
func reverseList(head *node.ListNode) *node.ListNode {
	var pre *node.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(head *node.ListNode) *node.ListNode {
	var tmp *node.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = tmp
		tmp = cur
		cur = next
	}
	return tmp
}

func main() {
	n5 := &node.ListNode{Val: 5, Next: nil}
	n4 := &node.ListNode{Val: 4, Next: n5}
	n3 := &node.ListNode{Val: 3, Next: n4}
	n2 := &node.ListNode{Val: 2, Next: n3}
	n1 := &node.ListNode{Val: 1, Next: n2}
	rel := reverseList(n1)
	p := list(rel)
	fmt.Println(p)
}

func list(head *node.ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
