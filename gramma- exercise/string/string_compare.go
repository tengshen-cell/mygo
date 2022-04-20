package main

import "fmt"

func main() {

	// ==
	s1 := "世界和平"
	s2 := "世界" + "和平"
	fmt.Println(s1 == s2)

	s1 = "Go"
	s2 = "C"
	fmt.Println(s1 != s2)

	s1 = "12345"
	s2 = "23456"
	fmt.Println(s1 < s2)
	fmt.Println(s1 <= s2)

	s1 = "12345"
	s2 = "123"
	fmt.Println(s1 > s2)
	fmt.Println(s1 >= s2)
}
