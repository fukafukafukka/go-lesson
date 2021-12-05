package main

import (
	"fmt"
	"time"
)

func goroutine_1(ch chan string) {
	for {
		ch <- "packet from fast network"
		time.Sleep(5 * time.Second)
	}
}

func goroutine_2(ch chan string) {
	for {
		ch <- "packet from slow network"
		time.Sleep(8 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine_1(c1)
	go goroutine_2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
