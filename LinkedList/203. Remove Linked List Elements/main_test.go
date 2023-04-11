package RemoveLinkedListElements

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name     string
		input    *ListNode
		val      int
		expected *ListNode
	}{
		{
			name:     "val is in the middle of list",
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			val:      3,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		},
		{
			name:     "val is at the beginning of the list",
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			val:      1,
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		},
		{
			name:     "val is at the end of the list",
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			val:      4,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name:     "empty list",
			input:    nil,
			val:      1,
			expected: nil,
		},
		{
			name:     "list has only one element",
			input:    &ListNode{Val: 1},
			val:      1,
			expected: nil,
		},
		{
			name:     "list has no matching element",
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			val:      4,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeElements(tc.input, tc.val)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func BenchmarkRemoveElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElements(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, 3)
	}
}

func BenchmarkRemoveElementsRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElementsRecursive(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, 3)
	}
}
