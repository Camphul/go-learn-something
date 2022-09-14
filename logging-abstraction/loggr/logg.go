package loggr

type Logg interface {
	Log(level Level, msg string)
	Info(msg string)
}
