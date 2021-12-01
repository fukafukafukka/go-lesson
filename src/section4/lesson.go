package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int
	for _, v := range l {
		if min == 0 {
			min = v
			continue
		}
		if v <= min {
			min = v
		}
	}
	fmt.Println(min)
}
