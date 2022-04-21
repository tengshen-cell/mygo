package main

import "fmt"

func main() {
	m := map[int]string{}
	fmt.Println(m)

	m1 := map[int][]string{
		1: []string{"val1_1", "val1_2"},
		3: []string{"val3_1", "val3_2", "val3_2"},
		7: []string{"val7_1"},
	}
	type Position struct {
		x float64
		y float64
	}

	m2 := map[Position]string{
		Position{29.935523, 52.568915}:  "school",
		Position{25.352594, 113.304361}: "shopping-mail",
		Position{73.224455, 111.804306}: "hospital",
	}

	fmt.Println(m1)
	fmt.Println(m2)

	m3 := map[Position]string{
		{29.935523, 52.568915}:  "school",
		{25.352594, 113.304361}: "shopping-mail",
		{73.224455, 111.804306}: "hospital",
	}
	fmt.Println(m3)
}
