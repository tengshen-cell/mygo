package main

import "fmt"

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	fmt.Printf("{")
	for k, v := range m {
		fmt.Printf("[%d,  %d]", k, v)
	}
	fmt.Printf("}\n")
	for k, _ := range m {
		fmt.Printf("[%d]", k)
	}
	fmt.Println()
	for k := range m {
		fmt.Printf("[%d]", k)
	}
}
