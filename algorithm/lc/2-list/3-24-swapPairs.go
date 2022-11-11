package main

import (
	node "gopl/algorithm/lc/common"
)

/**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
		1 -> 2 -> 3 -> 4
				|
				V
		2 -> 1 -> 4 -> 3
*/

func main() {

}

func swapPairs(head *node.ListNode) *node.ListNode {
	dummy := &node.ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		tmp := head.Next.Next
		head.Next.Next = head
		head.Next = tmp
		pre = head
		head = tmp
	}
	return dummy.Next
}
