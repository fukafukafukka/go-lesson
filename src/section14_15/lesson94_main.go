package main

import (
	// go mod init section14
	// go mod edit -replace config=./config
	// go get config
	"config"
	"log"

	// go mod edit -replace utils=./utils
	// ge get utils
	"utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	// fmt.Println(config.Config.ApiKey)
	// fmt.Println(config.Config.ApiSecret)
}
