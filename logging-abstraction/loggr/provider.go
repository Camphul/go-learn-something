package loggr

import "fmt"

var defaultLoggr Logg = nil

func DefaultLoggr() Logg {
	if defaultLoggr != nil {
		return defaultLoggr
	}
	fmt.Println("Creating default logger...")
	defaultLoggr = &LogLogg{}
	return defaultLoggr
}

func CreateFmtLoggr() Logg {
	return &FmtLogg{}
}
