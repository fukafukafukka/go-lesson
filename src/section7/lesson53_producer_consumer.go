package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	// Something
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		// 以下deferはconsumerメソッドが終わったタイミングで実行されるので、リソースリーク等につながるのでよろしくない。
		// defer wg.Done()
		// fmt.Println("process", i*1000)

		// 以下のように内部関数を作成して、その中にdeferを定義する方が良い。
		func() {
			defer wg.Done()
			fmt.Println("process", i*100)
		}()
	}

	fmt.Println("この標準出力がされているということは、mainスレッド側でclose(ch)してもらっているということ。※closeしてもらわないとfor文でchを待ち続けてしまう。")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)

	// mainスレッドで多少sleepしないと、consumer内の最後の処理の前にプロセスが終わってしまう。
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
