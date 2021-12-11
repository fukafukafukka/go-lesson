package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
	T         *T       `json:"T.omitempty"`
}

func main() {
	// Personの変数名と綴りさえ合っていれば、取得できる
	b := []byte(`{"name":"mike", "age":20, "nicknames":["a","b","c"]}`)
	var p Person

	// unmarshalの実行
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// marshalの実行
	v, _ := json.Marshal(p) // structのjson:"名前.型"で指定した変数名でjson生成される
	fmt.Println(string(v))
}
