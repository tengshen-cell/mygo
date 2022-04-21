package main

import "fmt"

func main() {
	var s []int
	s = append(s, 11)
	fmt.Println(len(s), cap(s))

	s = append(s, 12)
	fmt.Println(len(s), cap(s))

	s = append(s, 13)
	fmt.Println(len(s), cap(s))

	s = append(s, 14)
	fmt.Println(len(s), cap(s))

	s = append(s, 15)
	fmt.Println(len(s), cap(s))

	u := [...]int{11, 12, 13, 14, 15}
	fmt.Println("array:", u)
	s1 := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s1), cap(s1), s1)

	s1 = append(s1, 24)
	fmt.Println("after append 24, array:", u)
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s1), cap(s1), s1)

	s1 = append(s1, 25)
	fmt.Println("after append 25, array:", u)
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s1), cap(s1), s1)

	s1 = append(s1, 26)
	fmt.Println("after append 26, array:", u)
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s1), cap(s1), s1)

	s1[0] = 22
	fmt.Println("after reassign lst elem of slice, array:", u)
	fmt.Printf("after reassign lst elem of slice, slice(len=%d, cap=%d): %v\n", len(s1), cap(s1), s1)
}
