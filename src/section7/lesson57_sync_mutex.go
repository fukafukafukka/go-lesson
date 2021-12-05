package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func errorConCurrentMapWrites() {
	c := make(map[string]int)
	go func() {
		for i := 0; i < 10; i++ {
			c["key"] += 1
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c["key"] += 1
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c["key"])
}

func normalMapWrites() {
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c)
	fmt.Println(c.Value("Key"))
}

func main() {
	// 以下は同時にcに書き込みが生じるとエラー終了となる
	// errorConCurrentMapWrites()

	normalMapWrites()
}
