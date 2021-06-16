// bgf_log提供可配置的分层的日志系统，并可兼容golang官方包中的log.Logger
// bgf_log配置的搜索路径是:
// TODO:与运维确认下面的两个路径
// 1. $HOME/.bgf/log.yaml
// 2. 项目部署根目录/config/log.yaml
// 项目部署根目录可以通过getDeployRootPath(true)获取
// 如果上面的都没有，将日志输出到标准输出
package bgf_log

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/go-yaml/yaml"
)

// 日志类型
const (
	// 标准输出
	LOG_TYPE_STDOUT = "stdout"

	//输出到文件
	LOG_TYPE_FILE = "file"
)

const DEFAULT_CATEGORY = "*"

// LevelLogger接口提供按层级记录日志的方法，层级从Debug, Info, Warn, Error到Fatl
// 严重程度依次递增。开启某个级别的日志，会输出此级别以及严重程度级别更高的日志，
// 生产环境需要开启Warn级别，开发环境一般开启Debug级别
type LevelLogger interface {
	Log(level Level, msg string)
	Logf(level Level, format string, v ...interface{})
	Debug(msg string)
	Debugf(format string, v ...interface{})
	Info(msg string)
	Infof(format string, v ...interface{})
	Warn(msg string)
	Warnf(format string, v ...interface{})
	Error(msg string)
	Errorf(format string, v ...interface{})
	Fatl(msg string)
	Fatlf(format string, v ...interface{})
}

// 日志配置
type levelLoggerConfig struct {
	// Enable可以设置此日志是否生效
	Enable bool `yaml:"enable"`

	Categories []string `yaml:"categories"`

	// LogLevel目前支持debug, info, warn, error, fatl五种日志级别
	LogLevel string `yaml:"logLevel"`

	// LogType目前支持两种类型，stdout和file。分别是标准输出和文件
	LogType string `yaml:"logType"`

	// LogFile只有在LogType为file时才有效
	LogFile string `yaml:"logFile"`

	//和 官方log 库中的 flag 一致 Llongfile, Lshortfile等
	LogFLag []string `yaml:"logFLag"`
}

type levelLogger struct {
	Enable     bool
	Categories []string
	LogLevel   Level
	LogFile    *os.File
	l          *log.Logger
	mu         sync.Mutex // 参考 golang 官方log 库
	flag       int        //参考golang 官方log 库
}

type IContext interface {
	GetRequestId() string
}

func (l *levelLogger) Write(p []byte) (n int, err error) {
	return l.LogFile.Write(p)
}

var isEqualFunc = func(strInSlice, str string) bool {
	if strInSlice == DEFAULT_CATEGORY || strInSlice == str {
		return true
	}
	return false
}

var categoryLogLevelList = make(map[string]Level, 0)

func getCategoryLogLevel(category string) (logLevel Level, ok bool) {
	logLevel, ok = categoryLogLevelList[category]

	return
}

func SetLogLevelMap(levelMap map[string]Level) {
	categoryLogLevelList = levelMap
}

func (l *levelLogger) Log(level Level, category, msg string, callerSkip int) {
	logLevel, ok := getCategoryLogLevel(category)
	if !ok {
		logLevel = l.LogLevel
	}
	if !l.Enable || logLevel > level || !stringSliceContains(l.Categories, category, isEqualFunc) {
		return
	}

	var file string
	var line int

	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Lock()
		var ok bool
		_, file, line, ok = runtime.Caller(callerSkip)
		if !ok {
			file = "???"
			line = 0
		}
		if l.flag&log.Lshortfile != 0 {
			file = filepath.Base(file)
		}
		l.l.Printf("%s:%d [%s] [%s] %s", file, line, level.String(), category, msg)
		l.mu.Unlock()
	} else {
		l.l.Printf("[%s] [%s] %s", level.String(), category, msg)
	}

	if level == LEVEL_FATAL {
		//如果是Fatl，程序会退出，输出调用栈
		os.Exit(1)
	}
}

func (l *levelLogger) LogWithContext(ctx IContext, level Level, category, msg string, callerSkip int) {
	if !l.Enable || l.LogLevel > level || !stringSliceContains(l.Categories, category, isEqualFunc) {
		return
	}

	var file string
	var line int
	var requestId string
	if ctx != nil {
		requestId = ctx.GetRequestId()
	}
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Lock()
		var ok bool
		_, file, line, ok = runtime.Caller(callerSkip)
		if !ok {
			file = "???"
			line = 0
		}
		if l.flag&log.Lshortfile != 0 {
			file = filepath.Base(file)
		}
		if requestId != "" {
			l.l.Printf("%s:%d [%s] [%s] [%s] %s", file, line, level.String(), category, requestId, msg)
		} else {
			l.l.Printf("%s:%d [%s] [%s] %s", file, line, level.String(), category, msg)
		}
		l.mu.Unlock()
	} else {
		if requestId != "" {
			l.l.Printf("[%s] [%s] [%s] %s", level.String(), category, requestId, msg)
		} else {
			l.l.Printf("[%s] [%s] %s", level.String(), category, msg)
		}
	}

	if level == LEVEL_FATAL {
		//如果是Fatl，程序会退出，输出调用栈
		os.Exit(1)
	}
}

func (l *levelLogger) Logf(level Level, category, format string, v ...interface{}) {
	l.Log(level, category, fmt.Sprintf(format, v...), 4)
}

var levelLoggerList []*levelLogger

var logPrefix = ""

//var defaultLogFlag = log.Lmicroseconds | log.Llongfile
var defaultLogFlag = log.Ldate | log.Ltime | log.Lmicroseconds
var once sync.Once

type ConfigFilenameFinder func() (string, error)

var GetConfigFilename ConfigFilenameFinder = func() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	configFilename := fmt.Sprintf("%s/.bgf/log.yaml", u.HomeDir)
	if !fileExists(configFilename) {
		//configFilename = fmt.Sprintf("%s/bgf_log/log.yaml", getDeployRootPath(true))
		configFilename = fmt.Sprintf("%s/bgf_log/log.yaml", ".")
		fmt.Println(configFilename)
		if !fileExists(configFilename) {
			configFilename = ""
		}
	}

	return configFilename, nil
}

func GetLogger(category string) *Logger {
	once.Do(func() {
		err := initLevelLoggerList()
		if err != nil {
			panic(fmt.Sprintf("初始化日志失败 err:%s stack:%s", err, stackNoNewLine(0)))
		}
	})

	return &Logger{
		Category:       category,
		StdLoggerLevel: LEVEL_WARN,
	}
}

func initLevelLoggerList() error {
	configFilename, err := GetConfigFilename()
	if err != nil {
		return err
	}
	var c = make([]levelLoggerConfig, 0)
	if configFilename != "" {
		c, err = parseConfig(configFilename)
		if err != nil {
			panic(fmt.Sprintf("解析日志配置失败 err:%s stack:%s", err, stackNoNewLine(0)))
		}
	}
	return newLevelLoggerList(c...)
}

func parseConfig(configFilename string) ([]levelLoggerConfig, error) {
	config := make(map[string]levelLoggerConfig)
	b, err := ioutil.ReadFile(configFilename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}
	ret := make([]levelLoggerConfig, 0)
	for _, c := range config {
		ret = append(ret, c)
	}
	return ret, nil
}

func newLevelLoggerList(c ...levelLoggerConfig) error {
	loggerList := make([]*levelLogger, 0)
	//如果没有任何日志配置，将输出到stdout
	if len(c) == 0 {
		c = make([]levelLoggerConfig, 0, 1)
		var defaultLogLevel Level = LEVEL_DEBUG
		var defaultCategories = []string{DEFAULT_CATEGORY}
		c = append(c, levelLoggerConfig{
			Enable:     true,
			Categories: defaultCategories,
			LogLevel:   defaultLogLevel.String(),
			LogType:    LOG_TYPE_STDOUT,
			LogFLag:    nil,
		})
	}
	for _, logConfig := range c {
		// 测试和生产环境通过环境变量注入日志文件收集的日志级别
		logLevel := os.Getenv("log_level")
		if logLevel == "" {
			logLevel = logConfig.LogLevel
		}
		ll := &levelLogger{
			Enable:     logConfig.Enable,
			Categories: logConfig.Categories,
			LogLevel:   GetLogLevelFromString(logLevel),
		}
		fmt.Println(logConfig)
		if logConfig.LogType == LOG_TYPE_STDOUT {
			ll.LogFile = os.Stdout
		} else {
			f, err := os.OpenFile(logConfig.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				return err
			}

			ll.LogFile = f
		}

		//计算 log flag
		logFlag := defaultLogFlag
		if logConfig.LogFLag != nil {
			for _, tempFlag := range logConfig.LogFLag {
				switch tempFlag {
				case "Ldate":
					logFlag |= log.Ldate
				case "Ltime":
					logFlag |= log.Ltime
				case "Lmicroseconds":
					logFlag |= log.Lmicroseconds
				case "Lshortfile":
					logFlag |= log.Lshortfile
				case "Llongfile":
					logFlag |= log.Llongfile
				case "LUTC":
					logFlag |= log.LUTC
				case "LstdFlags":
					logFlag |= log.LstdFlags
				}
			}
		}

		ll.flag = logFlag
		ll.l = log.New(ll, logPrefix, defaultLogFlag)

		loggerList = append(loggerList, ll)
	}

	levelLoggerList = loggerList

	return nil
}

// Logger是日志类，提供了同时向多个目标（标准输出或者文件）
// 输出日志的功能
type Logger struct {
	Category string

	//如果使用的第三方库（包括标准库）中使用了标准的log.Logger
	//作为日志记录，可以将此字段注入，此字段实现了聚合输出日志
	//的功能
	StdLogger *log.Logger

	//StdLogger的日志级别
	StdLoggerLevel Level
}

// 实现io.Writer接口
func (l *Logger) Write(p []byte) (n int, err error) {
	n = len(p)
	err = nil
	for _, ll := range levelLoggerList {
		var stdLoggerLevel Level = LEVEL_WARN
		if l.StdLoggerLevel != 0 {
			stdLoggerLevel = l.StdLoggerLevel
		}
		ll.Log(stdLoggerLevel, l.Category, string(p), 4)
	}

	return n, err
}

func (l *Logger) Log(level Level, msg string, callerSkip int) {
	for _, ll := range levelLoggerList {
		ll.Log(level, l.Category, msg, callerSkip)
	}
}

func (l *Logger) LogWithContext(ctx IContext, level Level, msg string, callerSkip int) {
	for _, ll := range levelLoggerList {
		ll.LogWithContext(ctx, level, l.Category, msg, callerSkip)
	}
}

func (l *Logger) Logf(level Level, format string, v ...interface{}) {
	l.Log(level, fmt.Sprintf(format, v...), 4)
}

func (l *Logger) LogfWithContext(ctx IContext, level Level, format string, v ...interface{}) {
	l.LogWithContext(ctx, level, fmt.Sprintf(format, v...), 4)
}

func (l *Logger) Debug(msg string) {
	l.Log(LEVEL_DEBUG, msg, 3)
}

func (l *Logger) DebugWithContext(ctx IContext, msg string) {
	l.LogWithContext(ctx, LEVEL_DEBUG, msg, 3)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(LEVEL_DEBUG, format, v...)
}

func (l *Logger) DebugfWithContext(ctx IContext, format string, v ...interface{}) {
	l.LogfWithContext(ctx, LEVEL_DEBUG, format, v...)
}

func (l *Logger) Info(msg string) {
	l.Log(LEVEL_INFO, msg, 3)
}

func (l *Logger) InfoWithContext(ctx IContext, msg string) {
	l.LogWithContext(ctx, LEVEL_INFO, msg, 3)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(LEVEL_INFO, format, v...)
}

func (l *Logger) InfofWithContext(ctx IContext, format string, v ...interface{}) {
	l.LogfWithContext(ctx, LEVEL_INFO, format, v...)
}

func (l *Logger) Warn(msg string) {
	l.Log(LEVEL_WARN, msg, 3)
}

func (l *Logger) WarnWithContext(ctx IContext, msg string) {
	l.LogWithContext(ctx, LEVEL_WARN, msg, 3)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logf(LEVEL_WARN, format, v...)
}

func (l *Logger) WarnfWithContext(ctx IContext, format string, v ...interface{}) {
	l.LogfWithContext(ctx, LEVEL_WARN, format, v...)
}

func (l *Logger) Error(msg string) {
	l.Log(LEVEL_ERROR, msg, 3)
}

func (l *Logger) ErrorWithContext(ctx IContext, msg string) {
	l.LogWithContext(ctx, LEVEL_ERROR, msg, 3)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(LEVEL_ERROR, format, v...)
}

func (l *Logger) ErrorfWithContext(ctx IContext, format string, v ...interface{}) {
	l.LogfWithContext(ctx, LEVEL_ERROR, format, v...)
}

func (l *Logger) Fatal(msg string) {
	l.Log(LEVEL_FATAL, msg, 3)
}

func (l *Logger) FatalWithContext(ctx IContext, msg string) {
	l.LogWithContext(ctx, LEVEL_FATAL, msg, 3)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logf(LEVEL_FATAL, format, v...)
}

func (l *Logger) FatalfWithContext(ctx IContext, format string, v ...interface{}) {
	l.LogfWithContext(ctx, LEVEL_FATAL, format, v...)
}
