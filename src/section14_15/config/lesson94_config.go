package config

// lesson100_config.goとダブルためコメントアウト

// import (
// 	"log"
// 	"os"

// 	// go mod init config
// 	// go get gopkg.in/ini.v1
// 	"gopkg.in/ini.v1"
// )

// type ConfigList struct {
// 	ApiKey    string
// 	ApiSecret string
// 	LogFile   string
// }

// var Config ConfigList

// func init() {
// 	cfg, err := ini.Load("lesson94_config.ini")
// 	if err != nil {
// 		log.Printf("Failed to read file: %v", err)
// 		os.Exit(1)
// 	}
// 	Config = ConfigList{
// 		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
// 		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
// 		LogFile:   cfg.Section("gotrading").Key("log_file").String(),
// 	}
// }
