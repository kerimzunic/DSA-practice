package main

import "fmt"

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

func (arr *DynamicArray) len() int {
	return arr.capacity
}

func main() {
	testArr := NewDynamicArray()
	testArr2 := NewDynamicArray()
	fmt.Println("Capacity 1 is: ", (*testArr).len())
	fmt.Println("Capacity 2 is: ", testArr2.len())
	fmt.Println("Capacity 1 is: ", testArr.len())
}
