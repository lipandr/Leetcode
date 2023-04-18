package RemoveNthNodeFromEndOfList

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := dummy

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

// 1 2 3 4 5  n=2
