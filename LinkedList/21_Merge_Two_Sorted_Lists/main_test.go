package MergeTwoSortedLists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	// Define test cases
	tests := []struct {
		name   string
		list1  *ListNode
		list2  *ListNode
		expect *ListNode
	}{
		{
			name:   "Both lists are empty",
			list1:  nil,
			list2:  nil,
			expect: nil,
		},
		{
			name:   "One list is empty",
			list1:  nil,
			list2:  &ListNode{Val: 1},
			expect: &ListNode{Val: 1},
		},
		{
			name:  "Both lists have one node",
			list1: &ListNode{Val: 1},
			list2: &ListNode{Val: 2},
			expect: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
		},
		{
			name:  "Both lists have multiple nodes",
			list1: &ListNode{Val: 1, Next: &ListNode{Val: 3}},
			list2: &ListNode{Val: 2, Next: &ListNode{Val: 4}},
			expect: &ListNode{
				Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4}}},
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("unexpected result got=%v, expect=%v", got, tt.expect)
			}
		})
	}
}
