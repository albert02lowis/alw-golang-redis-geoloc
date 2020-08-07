package main

import (
	"log"
	"os"
)

//MyLogFile refers to the log file instance
var MyLogFile *os.File

//InitializeLogs Create a new file for simple logging.
func InitializeLogs() {
	if MyLogFile == nil {
		//create your file with desired read/write permissions
		f, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//set output of logs to f
		log.SetOutput(f)
		//test case
		log.Println("Initialized Logging.")
		MyLogFile = f
	}
}
