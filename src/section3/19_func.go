package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func doubleRetern(x int, y int) (int, int) {
	return x + y, x - y
}

func cal(price int, item int) (result int) {
	// 上記のように返り値である変数名を示すことができる。使いどきは返り値が複数ある場合
	result = price * item
	// result := price * item // 関数定義内の返り値に変数resultを定義しているので、以下のように「:」を使って最適できない。
	return result
	// return // returnで「result」を定義しないで良い
}

func main() {
	r := add(10, 20)
	fmt.Println(r)
	fmt.Println("-------------------------------")

	r1, r2 := doubleRetern(10, 20)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println("-------------------------------")

	r3 := cal(100, 2)
	fmt.Println(r3)
	fmt.Println("-------------------------------")

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
	fmt.Println("-------------------------------")

	// 以下のようにその場で実行もできる
	func(x int) {
		fmt.Println("inner func", x)
	}(1)
	fmt.Println("-------------------------------")
}
