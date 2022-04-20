package main

import "fmt"

type myInt int

const n myInt = 13
const m int = int(n) + 5

const e = 13

func main() {
	var a int = 5
	fmt.Println(a + int(n))

	fmt.Println(a + e)
}
