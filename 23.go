/*
23. Merge k Sorted Lists

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	list := lists[0]
	for i := 1; i < len(lists); i++ {
		list = merge2Lists(list, lists[i])
	}
	return list
}

func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	val1 := list1.Val
	val2 := list2.Val
	if val1 < val2 {
		next := merge2Lists(list1.Next, list2)
		list1.Next = next
		return list1
	} else {
		next := merge2Lists(list1, list2.Next)
		list2.Next = next
		return list2
	}
}

func main() {
}
