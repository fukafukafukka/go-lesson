package main

import "fmt"

// 引数のchanに矢印をつけることができて、わかりやすいので推奨。
func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func multiDouble(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func multiiQuadruple(second <-chan int, third chan<- int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multiDouble(first, second)
	go multiiQuadruple(second, third)

	for result := range third {
		fmt.Println(result)
	}
}
