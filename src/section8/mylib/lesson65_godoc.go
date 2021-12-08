/*
mylib is my special lib.
*/
package mylib

// Average returns the average of a series of numbers
func Average3(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}

// 以下よりgodocコマンドは終了したので、webサーバ上でgodocを見ることもできない。
// https://go.dev/doc/go1.12#godoc

// 以下ページを参考にGoDocについて知っておくと良い。
// https://future-architect.github.io/articles/20200521/
