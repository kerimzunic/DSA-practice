package main

import (
	"fmt"
	"DSA-practice/dynamicarray"
)

func main() {
	testArr := dynamicarray.NewDynamicArray()
	testArr2 := dynamicarray.NewDynamicArray()
	fmt.Println("Capacity 1 is: ", (*testArr).Size())
	fmt.Println("Capacity 2 is: ", testArr2.Size())
	fmt.Println("Capacity 1 is: ", testArr.Size())
}
