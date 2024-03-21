package define

const (
	LogLevelALL LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelInvalid
)

func BuildDefaultLevelDescMap() map[LogLevel]string {
	return map[LogLevel]string{
		LogLevelALL:     "ALL",
		LogLevelDebug:   "DEBUG",
		LogLevelInfo:    "INFO",
		LogLevelWarning: "WARNING",
		LogLevelError:   "ERROR",
	}
}
