package logger

import (
	"log"
	"os"
	"time"
)

var File *os.File

func Init() {
	var err error
	logFileName := "logs/" + time.Now().Format("01-02-2006 15-04-05") + ".logger"
	if _, err = os.Stat("./logs"); os.IsNotExist(err) {
		_ = os.MkdirAll("./logs", os.ModePerm)
	}
	File, err = os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to open log file: %v", err)
	}
	log.SetOutput(File)
}

func Close() {
	err := File.Close()
	if err != nil {
		log.Panicf("Unable to close log file: %v", err)
	}
}
