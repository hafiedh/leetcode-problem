package problem

import "testing"

func TestTotalMoney(t *testing.T) {
	// Test case 1: n = 4
	n := 4
	expected := 10
	if result := TotalMoney(n); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: n = 10
	n = 10
	expected = 37
	if result := TotalMoney(n); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 3: n = 20
	n = 20
	expected = 96
	if result := TotalMoney(n); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 4: n = 0
	n = 0
	expected = 0
	if result := TotalMoney(n); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
