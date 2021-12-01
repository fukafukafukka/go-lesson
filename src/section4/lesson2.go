package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kini":   90,
	}

	var sum int
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}
