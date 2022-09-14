package loggr

import "fmt"

var defaultLoggr Logg = nil

func DefaultLoggr() Logg {
	if defaultLoggr != nil {
		return defaultLoggr
	}
	fmt.Println("Creating default logger...")
	return &FmtLogg{}
}
