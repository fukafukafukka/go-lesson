package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func logger() {
	LoggingSettings("test.log")

	log.Println("logging")

	log.Printf("%T %v", "test", "test")

	log.Fatalln("error!")
	log.Fatalf("%T %v", "test", "test") // ここも実行されない
	fmt.Println("上記で処理が終わるので実行されない")
}

func useCase() {
	_, err := os.Open("nothing.txt")
	if err != nil {
		log.Fatal("Exit", err)
	}
}

func main() {
	// logger()

	useCase()
}
