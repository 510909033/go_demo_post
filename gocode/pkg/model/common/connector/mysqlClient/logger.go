package mysqlClient

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var gormLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Millisecond * 100, // 慢 SQL 阈值
		//LogLevel:      logger.Error, // Log level
		LogLevel: logger.Info, // Log level
		Colorful: false,       // 禁用彩色打印
	},
)
