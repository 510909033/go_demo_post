package bgf_log

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// stack返回格式化的调用栈，skip参数用来设置忽略调用栈中的栈帧数量
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// stackNoNewLine返回当前调用栈的信息，不含有换行符。
// 日志中输出调用栈信息可以用此方法，用于某些日志采集程序不支持多行的情况
func stackNoNewLine(skip int) string {
	return string(bytes.Replace(stack(skip), []byte("\n"), []byte("\t"), -1))
}

// source返回第n行代码
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function返回程序计数器当前地址
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is lready included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

func stringSliceContains(slice []string, str string, isEqualFunc func(strInSlice, str string) bool) bool {
	if isEqualFunc == nil {
		isEqualFunc = func(strInSlice, str string) bool {
			return strInSlice == str
		}
	}
	for _, s := range slice {
		if isEqualFunc(s, str) {
			return true
		}
	}
	return false
}

// fileExists检查某个文件路径是否存在
func fileExists(filename string) bool {
	if filename == "" {
		return false
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// getDeployRootPath得到部署的根目录
// evalSymlinks参数是是否需要软链，一般情况传入true
func getDeployRootPath(evalSymlinks bool) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if evalSymlinks {
		ex, err = filepath.EvalSymlinks(ex)
		if err != nil {
			panic(err)
		}
	}

	return filepath.Dir(ex)
}
