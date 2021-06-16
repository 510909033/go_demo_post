package logger

import (
	"io"
	"log"
	"os"
)

var Level string

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
)

type Client struct {
	logger *log.Logger
}

func GetLogger() *Client {
	return &Client{
		logger: loggerClient,
	}
}

var loggerClient *log.Logger

func init() {
	out, err := os.OpenFile("e:/pkg.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	multiOut := io.MultiWriter(out, os.Stdout)
	loggerClient = log.New(multiOut, "", log.LstdFlags|log.Lshortfile)
}

func (l *Client) Debugf(format string, v ...interface{}) {
	l.logger.Printf(Debug+" "+format, v...)
}
func (l *Client) Infof(format string, v ...interface{}) {
	l.logger.Printf(Info+" "+format, v...)
}
func (l *Client) Warnf(format string, v ...interface{}) {
	l.logger.Printf(Warn+" "+format, v...)
}
func (l *Client) Errorf(format string, v ...interface{}) {
	l.logger.Printf(Error+" "+format, v...)
}
