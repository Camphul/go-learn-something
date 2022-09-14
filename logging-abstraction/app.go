package main

import (
	"fmt"
	"log"
	"logging-abstraction/loggr"
	"time"
)

func main() {
	log.Println("default log pkg")
	fmt.Println("Running logging-abstraction")
	logg := loggr.DefaultLoggr()
	logg.Info("Hello world!")
	logSecondary()
	logg.Info("This is my second logging message.")
	logg.Log(loggr.Levels.Debug, "I hope this prints out a debug log message.")
	logg.Info("-- Round two --")
	logg.Info("Hello world!")
	go logSecondary()
	time.Sleep(time.Second * 2)
	logg.Info("This is my second logging message.")
	logg.Log(loggr.Levels.Debug, "I hope this prints out a debug log message.")
	time.Sleep(time.Second * 2)
}

func logSecondary() {
	logger := loggr.DefaultLoggr()
	logger.Info("This is a message inside the logSecondary function using the logger variable.")
	logger.Log(loggr.Levels.Debug, "And a debug level msg inside logSecondary")
}
