package main

import (
	"fmt"
	"unsafe"
)

func foo(arr [5]int) {

}

func main() {
	var arr1 [5]int
	// var arr2 [6]int
	// var arr3 [5]string

	foo(arr1)
	// foo(arr2)
	// foo(arr3)

	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度：", len(arr))
	fmt.Println("数组大小：", unsafe.Sizeof(arr))

	var arr2 = [6]int{
		11, 12, 13, 14, 15, 16,
	}

	var arr3 = [...]int{
		21, 22, 23,
	}
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%T\n", arr3)

	var arr4 = [...]int{
		99: 39,
	}
	fmt.Printf("%T\n", arr4)

}
