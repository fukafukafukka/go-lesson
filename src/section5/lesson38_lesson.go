package main

import "fmt"

func main() {
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	i = *p1 + *p2  // i = 100 + 200
	p2 = p1        // 300
	j = *p2 + i    // 300 + 300
	fmt.Println(j) // 600
}
