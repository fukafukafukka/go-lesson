package main

import (
	"fmt"

	"mylib"
	"mylib/under"
)

// GOPATHモードで紹介されているが、moduleモードが推奨になったため、こちらで対応する。
// https://go.dev/doc/tutorial/create-module
// cd mylib
// go mod init mylib
// cd ..
// go mod init section8
// go mod edit -replace mylib=./mylib
// go mod tidy
// go run .
func main() {
	s := []int{1, 2, 3, 4, 5}
	result := mylib.Average(s)
	fmt.Println(result)

	mylib.Say()

	under.Hello()
}
