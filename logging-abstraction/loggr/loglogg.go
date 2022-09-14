package loggr

import (
	"fmt"
	"time"
)

type FmtLogg struct{ Logg }

func (l *FmtLogg) Log(level Level, msg string) {
	fmt.Printf("[%d][%s] %s\n", time.Now().UnixMilli(), level.Prefix, msg)
}

func (logg *FmtLogg) Info(msg string) {
	logg.Log(Levels.Info, msg)
}

func (logg *FmtLogg) Warn(msg string) {
	logg.Log(Levels.Warn, msg)
}
