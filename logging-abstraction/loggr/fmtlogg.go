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
	logg.Log(Info, msg)
}
