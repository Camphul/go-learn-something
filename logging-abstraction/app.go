package main

import (
	"fmt"
	"logging-abstraction/loggr"
)

func main() {
	fmt.Println("Running logging-abstraction")
	logg := loggr.DefaultLoggr()
	logg.Info("Hello world!")
	logg.Info("This is my second logging message.")
	logg.Log(loggr.Debug, "I hope this prints out a debug log message.")
}
