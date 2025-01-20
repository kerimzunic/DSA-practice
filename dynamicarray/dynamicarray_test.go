package dynamicarray

import "testing"

func TestNewDynamicArray(t *testing.T) {
	dynamicArray := NewDynamicArray()
	if dynamicArray.Size() != 0 {
		t.Error("Initial size should be 10")
	}
}