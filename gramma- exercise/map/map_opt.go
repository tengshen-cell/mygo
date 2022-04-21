package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "value1"
	m[2] = "value2"
	m[3] = "value3"
	fmt.Println(m)

	m1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(len(m1))
	m1["key3"] = 3
	fmt.Println(len(m1))

	m2 := make(map[string]int)
	v := m2["key1"]
	fmt.Println(v)

	m3 := make(map[string]int)
	v, ok := m3["key1"]
	if !ok {
		fmt.Println("key1 not map")
	}

	m4 := make(map[string]int)
	_, ok = m4["key1"]
	if !ok {
		fmt.Println("value not map")
	}

	m5 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m5)
	delete(m5, "key2")
	fmt.Println(m5)

}
