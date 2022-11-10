package main

import (
	node "gopl/algorithm/lc/common"
)

/**
description:给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
*/

func removeElements(head *node.ListNode, val int) *node.ListNode {
	tmpHead := &node.ListNode{}
	tmpHead.Next = head
	cur := tmpHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return tmpHead.Next
}
