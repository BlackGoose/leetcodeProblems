package main

func reverseList(head *ListNode) *ListNode {
	return recure(head, nil)
}

func recure(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return tail
	}
	next := head.Next
	head.Next = tail
	return recure(next, head)
}
