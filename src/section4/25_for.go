package main

import "fmt"

func main() {
	// 基本的なfor文の書き方
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("--------------")

	// for文の省略記法
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// for {
	// fmt.Println("無限ループ")
	// }
}
