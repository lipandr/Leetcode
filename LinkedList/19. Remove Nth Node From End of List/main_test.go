package RemoveNthNodeFromEndOfList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// Test case 1: Removing the last node from a list with more than one element
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	n1 := 1
	expected1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	result1 := removeNthFromEnd(head1, n1)
	assert.Equal(t, expected1, result1)

	// Test case 2: Removing the first node from a list with more than one element
	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	n2 := 4
	expected2 := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}
	result2 := removeNthFromEnd(head2, n2)
	assert.Equal(t, expected2, result2)

	// Test case 3: Removing the only node from a list with one element
	head3 := &ListNode{1, nil}
	n3 := 1
	expected3 := (*ListNode)(nil)
	result3 := removeNthFromEnd(head3, n3)
	assert.Equal(t, expected3, result3)
}
