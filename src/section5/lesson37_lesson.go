package main

import "fmt"

func main() {
	var i int = 10
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j)
}
