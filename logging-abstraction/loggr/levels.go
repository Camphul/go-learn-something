package loggr

type Level struct {
	Prefix string
}

type logLevels struct {
	Info  Level
	Warn  Level
	Debug Level
}

var Levels = logLevels{
	Info:  Level{"INFO"},
	Debug: Level{"DEBUG"},
	Warn:  Level{"WARN"},
}
