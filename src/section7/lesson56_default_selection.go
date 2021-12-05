package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	// 以下OuterLoop使えばfor文終わった後もプロセス継続される
OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
		// ここにbreak置くとfor文が1周目で終わる
	}
	fmt.Println("This is reachable for OuterLoop")

	// 以下はfor文終了するとプロセスも終了してしまう書き方
	tick2 := time.Tick(100 * time.Millisecond)
	boom2 := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick2:
			fmt.Println("tick.")
		case <-boom2:
			fmt.Println("BOOM!")
			return // breakだとfor文自体は終わらないので無限ループとなる
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
		// ここにbreak置くとfor文が1周目で終わる
	}
	fmt.Println("This is unreachable for return")
}
