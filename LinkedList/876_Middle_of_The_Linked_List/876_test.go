package middleNode

import (
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}},

			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}},

		{&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4, nil}}}},

			&ListNode{3,
				&ListNode{4, nil}}},

		{&ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}},

			&ListNode{2, &ListNode{3, nil}}},

		{&ListNode{1, nil},

			&ListNode{1, nil}},
	}

	for _, test := range tests {
		result := middleNode(test.head)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected output %v but got %v", test.head, test.expected, result)
		}
	}
}
