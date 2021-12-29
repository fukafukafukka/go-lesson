package main

import (
	// go mod init section14_15
	// go mod edit -replace config=./config
	// go get config
	"config"
	"fmt"

	// go mod edit -replace utils=./utils
	// go get utils
	"utils"

	// go mod edit -replace bitflyer=./bitflyer
	// go get bitflyer
	"bitflyer"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
