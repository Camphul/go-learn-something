package loggr

type Level struct {
	Prefix string
}

var Info = Level{"INFO"}
var Debug = Level{"DEBUG"}
