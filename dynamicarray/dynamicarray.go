package dynamicarray

// This will only work for integers for now
type DynamicArray struct {
	data     *[10]int
	capacity int
	size     int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data:     new([10]int),
		capacity: 10,
		size:     0,
	}
}

func NewDynamicArrayWithCapacity(capacity int) *DynamicArray {
	return &DynamicArray{
		data:     new([10]int),
		capacity: capacity,
		size:     0,
	}
}

func NewDynamicArrayFromStaticArray(static_array []int) *DynamicArray {
	return &DynamicArray{
		data:     new([10]int),
		capacity: 0,
		size:     0,
	}
}

func (arr *DynamicArray) resize(elements int) int {
	return arr.size
}

func (arr *DynamicArray) GetSize() int {
	return arr.size
}

func (arr *DynamicArray) GetCapacity() int {
	return arr.capacity
}

func (arr *DynamicArray) Add() int {
	return arr.size
}

func (arr *DynamicArray) Remove() int {
	return arr.size
}

func (arr *DynamicArray) Get(i int) int {
	if i < 0 && i > arr.size {
		panic("Out of bounds access!")
	}
	return arr.data[i]
}
