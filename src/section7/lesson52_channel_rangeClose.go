package main

import "fmt"

func goroutine_1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // make(chan int, len(S))とも宣言できる
	go goroutine_1(s, c)
	for i := range c {
		// goroutine_1メソッドでsumされる度にmainスレッドのcに値が入る
		fmt.Println(i) // 1,3,6,10,15
	}
}
