package bgf_log

import "strings"

// 日志层级类型
type Level int

func (ll Level) String() string {
	if ll < LEVEL_DEBUG || ll > LEVEL_FATAL {
		return LEVEL_NAME_UNKNOWN
	}

	return levelName[ll]
}

// GetLogLevelFromString 将日志层级的字符串表示转化为LogLevel类型的表示
func GetLogLevelFromString(level string) Level {
	var ll Level = LEVEL_UNKNOWN
	level = strings.ToLower(level)
	switch level {
	case "debug":
		ll = LEVEL_DEBUG
	case "info":
		ll = LEVEL_INFO
	case "warn":
		ll = LEVEL_WARN
	case "error":
		ll = LEVEL_ERROR
	case "fatal":
		ll = LEVEL_FATAL
	}

	return ll
}

// 日志级别，重要性从低到高一次递增，LEVEL_UNKNOWN不是一种可用类型
const (
	LEVEL_UNKNOWN = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

var LEVEL_NAME_UNKNOWN = "UNKNOWN"

var levelName = [...]string{
	LEVEL_NAME_UNKNOWN,
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}
