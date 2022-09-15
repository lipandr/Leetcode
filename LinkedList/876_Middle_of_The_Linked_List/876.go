package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	var array []*ListNode
	length := 0
	for head != nil {
		array = append(array, head)
		head = head.Next
		length++
	}
	return array[length/2]
}
