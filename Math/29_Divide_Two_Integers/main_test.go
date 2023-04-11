package DivideTwoIntegers

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	// Test cases where divisor is 0
	if divide(1, 0) != math.MaxInt32 {
		t.Errorf("divide(1, 0) = %d; want %d", divide(1, 0), math.MaxInt32)
	}
	if divide(-1, 0) != math.MaxInt32 {
		t.Errorf("divide(-1, 0) = %d; want %d", divide(-1, 0), math.MaxInt32)
	}

	// Test cases where dividend is 0
	if divide(0, 1) != 0 {
		t.Errorf("divide(0, 1) = %d; want 0", divide(0, 1))
	}
	if divide(0, -1) != 0 {
		t.Errorf("divide(0, -1) = %d; want 0", divide(0, -1))
	}

	// Test cases with valid inputs
	if divide(10, 3) != 3 {
		t.Errorf("divide(10, 3) = %d; want 3", divide(10, 3))
	}
	if divide(7, -3) != -2 {
		t.Errorf("divide(7, -3) = %d; want -2", divide(7, -3))
	}
	if divide(-2147483648, -1) != 2147483647 {
		t.Errorf("divide(-2147483648, -1) = %d; want %d", divide(-2147483648, -1), 2147483647)
	}
}
