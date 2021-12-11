package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"), // デフォルト値
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

// https://pkg.go.dev/github.com/go-ini/ini#section-readme
// go get gopkg.in/ini.v1
func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
