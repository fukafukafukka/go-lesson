package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 100
	fmt.Println(len(ch))

	ch <- 200
	fmt.Println(len(ch))

	// 以下は、channelのサイズ2に対して、3つ目になるのでエラーとなる
	// ch <- 300
	// fmt.Println(len(ch))

	x := <-ch
	fmt.Println(x)

	// 上記で1つ取り出した場合は、新たに追加できる
	ch <- 300
	fmt.Println(len(ch))

	close(ch)
	// 上記close処理をしないと、範囲外のchの値も出力しようとしてエラーとなる。
	for c := range ch {
		fmt.Println(c)
	}
}
