package dynamicarray

type DynamicArray struct {
	data     *[10]int
	capacity int
	size     int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data:     new([10]int),
		capacity: 0,
		size:     0,
	}
}

func (arr *DynamicArray) Size() int {
	return arr.size
}
