package main

import (
	// go mod init section14_15
	// go mod edit -replace config=./config
	// go get config
	"config"

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
	// fmt.Println(apiClient.GetBalance())

	// ticker, _ := apiClient.GetTicker("BTC_USD")
	// fmt.Println(ticker)
	// fmt.Println(ticker.GetMidPrice())
	// fmt.Println(ticker.DateTime())
	// fmt.Println(ticker.TruncateDateTime(time.Hour))

	tickerChannel := make(chan bitflyer.Ticker)
	apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
}
