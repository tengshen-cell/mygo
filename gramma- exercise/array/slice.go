package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(nums))

	nums = append(nums, 7)
	fmt.Println(len(nums))

	sl1 := make([]byte, 6, 10)
	sl2 := make([]byte, 6)
	fmt.Println(len(sl1))
	fmt.Println(len(sl2))

	// 数组的切片化
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl3 := arr[3:7:9]
	fmt.Println(sl3)

	sl3[0] += 10
	fmt.Println("arr[3]=", arr[3])
}
