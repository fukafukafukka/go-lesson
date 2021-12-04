package main

import "fmt"

func do_all(i interface{}) {}

func do_int(i interface{}) {
	ii := i.(int) // タイプアサーション（interfaceを型指定）してあげないといけない
	ii *= 2       // i = i * 2
	fmt.Println(ii)
}

func do_string(i interface{}) {
	ss := i.(string)
	fmt.Println(ss + "!")
}

func do_switch_type(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	do_all("Mike")
	do_all(true)
	do_all(10)

	fmt.Println("----------")

	do_int(10)

	fmt.Println("----------")

	// 以下のようにinterfaceに入れた場合でもメソッド先でint型に変換してもらう必要がある
	var i interface{} = 10
	do_int(i)

	fmt.Println("----------")

	do_string("Mike")

	fmt.Println("----------")

	do_switch_type(10)
	do_switch_type("Mike")
	do_switch_type(true)

	fmt.Println("----------")

	// ちなみに以下はタイプコンバージョン（キャスト）なので
	// タイプアサーションとは異なることに注意
	var x int = 10
	xx := float64(x)
	fmt.Println(xx)
}
