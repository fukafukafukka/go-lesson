package utils

// go mod init utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ref) section4/29_log.go
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}
