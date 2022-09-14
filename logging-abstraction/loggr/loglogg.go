package loggr

import (
	"log"
)

type LogLogg struct{ Logg }

func (l *LogLogg) Log(level Level, msg string) {
	log.Printf("[%s] %s\n", level.Prefix, msg)
}

func (logg *LogLogg) Info(msg string) {
	logg.Log(Levels.Info, msg)
}

func (logg *LogLogg) Warn(msg string) {
	logg.Log(Levels.Warn, msg)
}
