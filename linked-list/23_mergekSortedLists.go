package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func pop(lists *[]*ListNode) *ListNode {
	if len(*lists) == 0 {
		return nil
	}
	returnNode := (*lists)[0]
	*lists = (*lists)[1:]
	return returnNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var mergeList []*ListNode
	i := 0
	j := len(lists) - 1
	for i <= j {
		if j == i {
			mergeList = append(mergeList, lists[i])
		} else {
			mergeList = append(mergeList, mergeTwoLists(lists[i], lists[j]))
		}
		i++
		j--
	}

	for len(mergeList) != 1 {
		node1 := pop(&mergeList)
		node2 := pop(&mergeList)
		mergeList = append(mergeList, mergeTwoLists(node1, node2))
	}
	return pop(&mergeList)
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}
// 	if list1.Val < list2.Val {
// 		return &ListNode{list1.Val, mergeTwoLists(list1.Next, list2)}
// 	}
// 	return &ListNode{list2.Val, mergeTwoLists(list1, list2.Next)}
// }
