package gologger

const (
	//LogError is used to log any sort of error, both internal and external
	LogError = 0

	//LogCritical is used to log any large issues/errors
	LogCritical = 1

	//LogTraffic logs web traffic
	LogTraffic = 2

	//SystemLogIfCreateFail will log to stdout and LogError if the creation of the Logger fails
	SystemLogIfCreateFail = 0
)

var (
	//Error is used to log any sort of error, both internal and external
	Error = LogType{0}

	//Critical is used to log any large issues/errors
	Critical = LogType{1}

	//Traffic logs web traffic
	Traffic = LogType{2}
)

type LogType struct {
	int
}
