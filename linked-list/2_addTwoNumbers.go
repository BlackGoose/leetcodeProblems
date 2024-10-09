package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode = nil

	overall := 0

	var last *ListNode = nil
	var temp *ListNode = nil

	for l1 != nil || l2 != nil || overall > 0 {
		sum := overall
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		overall = sum / 10
		val := sum % 10

		if last == nil {
			last = &ListNode{Val: val}
			result = last
		} else {
			temp = &ListNode{Val: val}
			last.Next = temp
			last = temp
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return result
}
