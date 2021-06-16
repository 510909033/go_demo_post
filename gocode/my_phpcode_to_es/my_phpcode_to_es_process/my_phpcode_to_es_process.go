package my_phpcode_to_es_process

import (
	"baotian0506.com/39_config/gocode/my_phpcode_to_es/bgf_log"
	"encoding/json"
	"fmt"

	"io/ioutil"

	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	LOG_SAVE_DIR = "e:/logs/phpcode"
	LOG_PREFIX   = "PHP_CODE_"
)

func exception(err error) {
	if err != nil {
		panic(err)
	}
}

type Data struct {
	IsDir    bool
	FileName string
	Content  string
}

var f *os.File
var filename string
var mu sync.Mutex
var walkCh = make(chan *Data, 0)

//var logger = log.New(os.Stdout,"", log.LstdFlags|log.Lshortfile)
var logger = bgf_log.GetLogger("my_phpcode_to_es_process")

var walkDir = `E:\code_180622\all\pregnancy\scripts`

type Service struct {
}

func NewService() *Service {

	return &Service{}
}

// 将 walkDir 目录的所有文件夹和文件内容 放到 LOG_SAVE_DIR 里
func (s *Service) Start() {

	defer s.closeF()
	logger.Infof("Start, walkDir=%s, LOG_SAVE_DIR=%s\n", walkDir, LOG_SAVE_DIR)

	go s.saveToLog()

	err := filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		d := &Data{}
		d.IsDir = info.IsDir()
		d.FileName = path
		if !d.IsDir {
			logger.Infof("walk one filename=%s\n", d.FileName)
			b, err := ioutil.ReadFile(d.FileName)
			exception(err)
			d.Content = string(b)
		}
		s.sendData(d)
		return nil
	})
	logger.Info("Start over\n")
	exception(err)
	time.Sleep(time.Second)
}

func (s *Service) sendData(d *Data) {
	walkCh <- d
}

func (s *Service) getLogFPtr() *os.File {
	mu.Lock()
	defer mu.Unlock()

	//newFilename :=fmt.Sprintf("%d_%d", time.Now().Year(), time.Now().YearDay())
	newFilename := fmt.Sprintf("%s_%d", time.Now().Format("2006-01-02_15"), time.Now().Minute())
	newFilename = fmt.Sprintf("%s/%s%s.log", LOG_SAVE_DIR, LOG_PREFIX, newFilename)
	if newFilename == filename && f != nil {
		return f
	}
	s.closeF()

	filename = newFilename
	logger.Infof("log filename=%s\n", filename)
	var err error
	f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	exception(err)
	return f
}

func (s *Service) closeF() {
	if f != nil {
		exception(f.Close())
	}
}

func (s *Service) saveToLog() {
	for data := range walkCh {
		if data == nil {
			panic("data is nil")
		}

		b, err := json.Marshal(data)
		exception(err)
		_, err = s.getLogFPtr().WriteString(string(b) + "\n")

		exception(err)
		//time.Sleep(time.Millisecond*1000)
	}
}
