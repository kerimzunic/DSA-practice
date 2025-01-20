package dynamicarray

import "testing"

func TestNewDynamicArray(t *testing.T) {
	arr := NewDynamicArray()

	if arr.GetSize() != 0 {
		t.Errorf("Expected size 0, but got %d", arr.GetSize())
	}

	if arr.GetCapacity() != 10 {
		t.Errorf("Expected capacity 10, but got %d", arr.GetCapacity())
	}
}
