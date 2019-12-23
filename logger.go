package main

import (
	"log"
	"os"
)

func loggerPrint(str string) {
	logFile, err := os.OpenFile("debug.log", os.O_APPEND | os.O_RDWR | os.O_CREATE , 0666)
	if err != nil {
		log.Println("Create Debug Log File Failed!")
		return
	}
	defer logFile.Close()
	logger := log.New(logFile, "[INFO]", log.Ldate | log.Ltime | log.Lshortfile)
	logger.Println(str)
}
