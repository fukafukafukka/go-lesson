package main

import (
	"context"
	"fmt"
	"time"

	// サードパーティーライブラリの注入の仕方
	// go init section11
	// go get golang.org/x/sync/semaphore
	// 以下のようにソースコード中のimport箇所にパッケージを記述する
	// go mod tidy
	// 上記実行することにより、go.modファイル内にダウンロードしたsemaphoreを使うよう依存関係が記載される。
	// go run lesson76_semaphore.go
	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// isAcquire := s.TryAcquire(1)
	// if !isAcquire {
	// fmt.Println("Could not get lock")
	// return
	// }

	// ここでsemaphoreに設定した数のプロセスだけにロックする
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)

	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

// https://pkg.go.dev/golang.org/x/sync/semaphore
// サードパーティーのパッケージ
func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
