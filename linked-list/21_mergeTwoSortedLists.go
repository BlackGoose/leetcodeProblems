package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		return &ListNode{list1.Val, mergeTwoLists(list1.Next, list2)}
	}
	return &ListNode{list2.Val, mergeTwoLists(list1, list2.Next)}
}
