package main

import (
	"fmt"
	"sync"
	"time"
)

func normal_1(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal_2(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine_1(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine_2(s string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
	wg.Done()
}

// goroutineは別スレッドで実行されるため、メモリの値の引き渡しができない -> 引き渡しさせるには次のchanelを使う
func main() {
	normal_1("world")
	normal_2("hello")

	fmt.Println("----------")

	// 並列処理はgoとつけるだけ
	go goroutine_1("world")
	// normal_2メソッドのfor文内でsleepを実行しなかった場合、goroutineが実行される前にプログラムが終わる。
	normal_2("hello")

	fmt.Println("----------")

	// WaitGroupをgoroutine内メソッドに渡しつつ、mainスレッドの最後でWaitさせれば、sleepさせずともgoroutineが実行される。
	var wg sync.WaitGroup
	wg.Add(1) // 1つWaitGroupを追加したら必ず1つDoneしないといけない
	go goroutine_2("world", &wg)
	normal_2("hello")
	wg.Wait()
}
