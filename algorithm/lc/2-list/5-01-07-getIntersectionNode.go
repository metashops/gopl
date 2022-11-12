package main

import (
	node "gopl/algorithm/lc/common"
)

// 面试题 02.07. 链表相交

// func main() {
// 	n7 := &node.ListNode{Val: 7, Next: nil}
// 	n6 := &node.ListNode{Val: 6, Next: n7}
// 	n5 := &node.ListNode{Val: 5, Next: n6}
// 	n4 := &node.ListNode{Val: 4, Next: n5}
// 	n3 := &node.ListNode{Val: 3, Next: n4}
//
// 	n9 := &node.ListNode{Val: 9, Next: n5}
// 	n8 := &node.ListNode{Val: 8, Next: n9}
//
// 	b1 := &node.ListNode{Val: 2, Next: n8}
// 	a1 := &node.ListNode{Val: 1, Next: n3}
// 	rel := getIntersectionNode(a1, b1)
// 	fmt.Println(rel.Next.Val)
// }

// 双指针
func getIntersectionNode(headA, headB *node.ListNode) *node.ListNode {
	l1, l2 := headA, headB
	// 指针相等会相等
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}
