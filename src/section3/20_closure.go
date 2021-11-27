package main

import "fmt"

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	// 以下をメソッド化した。
	// x := 0
	// increment := func() int {
	// 	x++
	// 	return x
	// }
	// fmt.Println(increment()) //1
	// fmt.Println("-------------------------------")
	// fmt.Println(increment()) //2
	// fmt.Println("-------------------------------")
	// fmt.Println(increment()) //3
	// fmt.Println("-------------------------------")

	counter := incrementGenerator()
	fmt.Println(counter()) //1
	fmt.Println("-------------------------------")
	fmt.Println(counter()) //2
	fmt.Println("-------------------------------")
	fmt.Println(counter()) //3
	fmt.Println("-------------------------------")
	fmt.Println(counter()) //4
	fmt.Println("-------------------------------")

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println("-------------------------------")

	c2 := circleArea(3)
	fmt.Println(c2(2))
	fmt.Println("-------------------------------")
}
