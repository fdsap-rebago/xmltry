package middleware

import (
	"log"
	"os"
	"time"
)

func SetLogs() {
	logFileName := time.Now().Format("2006-01-02") + ".log"
	// create a new log with file name xxx or more the existing one
	f, err := os.OpenFile("insta-"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// set default log output to the 'new' file
	log.SetOutput(f)
	log.Println("This is a test log entry")
	defer f.Close()
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
