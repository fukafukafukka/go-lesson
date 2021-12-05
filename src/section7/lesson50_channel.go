package main

import "fmt"

func goroutine_1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine_2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// channelを使えば、別スレッドの値をmainスレッドで使うことができる
func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	go goroutine_1(s, c)
	go goroutine_2(s, c)

	// 以下は次のイメージで捉えておくとわかりやすい。x <- c <- sum
	x := <-c // chanelが値を受信するのを待つので、WaitGroup使う必要がない。
	fmt.Println(x)

	y := <-c // xに値を入れたらcの中身はなくなる
	fmt.Println(y)
}
