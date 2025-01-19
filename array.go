package main

import "fmt"

type DynamicArray struct {
	Capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        Capacity:  capacity,
    }
}

func main () {
	testArr := NewDynamicArray(2)
	fmt.Print(testArr)
}
