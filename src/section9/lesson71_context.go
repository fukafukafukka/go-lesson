package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	// goroutineの処理が終わらない時を想定して、処理待ち時間を以下のコンテキストというもので設定する
	ctx := context.Background() // contextの内容が具体的に決まっていない時はとりあえずcontext.TODO()をctxとしておいて開発するのも良い。
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	// コンテキストはgoroutineに渡すだけで良い。
	go longProcess(ctx, ch)
	// cancel()// タイムアウトでなくてこのようにキャンセルメソッドを使ってもgoroutineを終了させることができる

CTXLOOP:
	for {
		select {
		// contextのタイムアウト発生時
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		// contextのタイムアウトより前に、チャネルchに値が入った時
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}

	fmt.Println("###############")
}
