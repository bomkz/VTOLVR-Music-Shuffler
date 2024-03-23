package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func tickTime() {
	for {
		tick <- true
		time.Sleep(100 * time.Millisecond)
	}
}

func readLog() {

	// Start log reading ticker
	go tickTime()

	readlog := false

	// Reads log every tick.
	for {
		<-tick

		// Reads log line by line
		logFile := readLogFile()
		scanLog := bufio.NewScanner(strings.NewReader(logFile))
		var logLinesTmp []string

		// Appends each log line to a variable.
		for scanLog.Scan() {
			logLinesTmp = append(logLinesTmp, scanLog.Text())
		}

		// Only takes action on log lines that are new, and ignores old ones.
		for y, x := range logLinesTmp {
			if y > (len(logLines) - 1) {
				if readlog {
					// Handles new log lines.
					logHandler(x)
				}

			}
		}
		logLines = logLinesTmp

		// Set to true after the first for loop passes
		readlog = true
	}
}

// Runs the appropriate function depending on the log line contents.
func logHandler(newline string) bool {
	if strings.Contains(newline, "FlightLogger:") && strings.Contains(newline, "has spawned.") {

		fmt.Println("Splash 1, bearing 0")
		bwa()
	}
	return false
}

// reads the log file.
func readLogFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	file, err := os.ReadFile(home + "\\AppData\\LocalLow\\Boundless Dynamics, LLC\\VTOLVR\\Player.log")
	if err != nil {
		log.Panic(err)
	}

	return string(file)
}
