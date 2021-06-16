package bgf_log

import (
	"io/ioutil"
	"os"
	"testing"
)

type testData struct {
	in   string
	want Level
}

var testDatas = []testData{
	{"unknow", 0},
	{"debug", 1},
	{"info", 2},
	{"warn", 3},
	{"error", 4},
	{"fatal", 5},
}

var logConfig = `
testlog:
  enable: true
  categories:
    - "*"
  logLevel: debug
  logType: stdout
  logFLag:
    - Ltime
    - Llongfile
`
var logConfigFileName string

func init() {
	f, err := ioutil.TempFile("", "bgf_log_test_")
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(logConfig)
	if err != nil {
		panic(err)
	}
	logConfigFileName = f.Name()
}

func TestNewLevelLoggerList(t *testing.T) {
	levelLogConfigs, err := parseConfig(logConfigFileName)
	if err != nil {
		t.Fatalf("ParseConfig fail, err: %v", err)
	}
	err = os.Remove(logConfigFileName)
	if err != nil {
		t.Errorf("Remove tempLogConfigFile fail, path: %s, err: %v", logConfigFileName, err)
	}
	newLevelLoggerList(levelLogConfigs...)
	if len(levelLoggerList) == 0 {
		t.Errorf("NewLevelLoggerList fail.")
	}
	levelLoggerList[0].Log(LEVEL_DEBUG, "*", "bgf_log_test", 2)
}

func TestDebug(t *testing.T) {
	t.Parallel()
	logger := GetLogger("debug_log")
	logger.Debug("This is debug_log.")
}

func TestDebugf(t *testing.T) {
	t.Parallel()
	logger := GetLogger("debugf_log")
	logger.Debugf("This is %s.", "debugf_log")
}

func TestInfo(t *testing.T) {
	t.Parallel()
	logger := GetLogger("info_log")
	logger.Info("This is info_log.")
}

func TestInfof(t *testing.T) {
	t.Parallel()
	logger := GetLogger("infof_log")
	logger.Infof("This is %s.", "infof_log")
}

func TestWarn(t *testing.T) {
	t.Parallel()
	logger := GetLogger("warn_log")
	logger.Warn("This is warn_log.")
}

func TestWarnf(t *testing.T) {
	t.Parallel()
	logger := GetLogger("warnf_log")
	logger.Warnf("This is %s.", "warnf_log")
}

func TestError(t *testing.T) {
	t.Parallel()
	logger := GetLogger("error_log")
	logger.Error("This is error_log.")
}

func TestErrorf(t *testing.T) {
	t.Parallel()
	logger := GetLogger("errorf_log")
	logger.Errorf("This is %s.", "errorf_log")
}

func TestLogf(t *testing.T) {
	t.Parallel()
	logger := GetLogger("logf_log")
	logLevel := GetLogLevelFromString("info")
	logger.Logf(logLevel, "This is %s", "logf_log")
}

func TestWrite(t *testing.T) {
	t.Parallel()
	logger := GetLogger("mylog")
	logger.Write([]byte("This is mylog."))
}

func TestGetLogLevelFromString(t *testing.T) {
	for _, v := range testDatas {
		res := GetLogLevelFromString(v.in)
		if res != v.want {
			t.Errorf("Failed to GetLogLevelFromString, want[%v], get[%v]", v.want, res)
		}
	}
}

func TestLevelLogLogf(t *testing.T) {
	err := newLevelLoggerList()
	if err != nil {
		t.Errorf("Failed to newLevelLoggerList, err: %v", err)
	}
	ll := levelLoggerList[0]
	logLevel := GetLogLevelFromString("info")
	ll.Logf(logLevel, "mylog", "Invoke levelLogLogf %s", "just test.")
}

type testContext struct {
}

func (tc *testContext) GetRequestId() string {
	return "trace-id-8888-998877"
}

func TestLogWithContext(t *testing.T) {
	logger := GetLogger("withContextLogTest")
	ctx := &testContext{}
	logger.DebugWithContext(ctx, "trace test debug")
	logger.DebugfWithContext(ctx, "trace test debugf")
	logger.InfoWithContext(ctx, "trace test info")
	logger.InfofWithContext(ctx, "trace test infof")
	logger.ErrorWithContext(ctx, "trace test error")
	logger.ErrorfWithContext(ctx, "trace test errorf by %s", "babytree")
	//logger.FatalWithContext(ctx, "trace test fatal")
	//logger.FatalfWithContext(ctx, "trace test fatalf")
}
