package conver

import (
	"testing"
)

func TestBytesToString(t *testing.T) {
	// Test case 1: Empty byte slice
	input1 := []byte{}
	expectedOutput1 := ""
	output1 := BytesToString(input1)
	if output1 != expectedOutput1 {
		t.Errorf("BytesToString(%v) = %q, expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: Non-empty byte slice
	input2 := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	expectedOutput2 := "Hello, World!"
	output2 := BytesToString(input2)
	if output2 != expectedOutput2 {
		t.Errorf("BytesToString(%v) = %q, expected %q", input2, output2, expectedOutput2)
	}
}

func TestStringToBytes(t *testing.T) {
	// Test case 1: Empty string
	input1 := ""
	expectedOutput1 := []byte{}
	output1 := StringToBytes(input1)
	if !isEqualSlice(output1, expectedOutput1) {
		t.Errorf("StringToBytes(%q) = %v, expected %v", input1, output1, expectedOutput1)
	}

	// Test case 2: Non-empty string
	input2 := "Hello, World!"
	expectedOutput2 := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	output2 := StringToBytes(input2)
	if !isEqualSlice(output2, expectedOutput2) {
		t.Errorf("StringToBytes(%q) = %v, expected %v", input2, output2, expectedOutput2)
	}
}

// Helper function to compare two byte slices
func isEqualSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
