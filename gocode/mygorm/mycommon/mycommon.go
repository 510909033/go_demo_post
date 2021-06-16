package mycommon

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"runtime"
	"time"
)

var db *gorm.DB
var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second,  // 慢 SQL 阈值
		LogLevel:      logger.Error, // Log level
		Colorful:      false,        // 禁用彩色打印
	},
)

func init() {
	//db, err := gorm.Open("mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//db.Debug()

	//defer db.Close()

	sqlDB, _ := db.DB()
	_ = sqlDB

	//sqlDB.Close()
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetConnMaxLifetime(time.Second)

	go func() {
		for {
			fmt.Printf("%+v\n", sqlDB.Stats())
			time.Sleep(time.Second)
		}
	}()

}

func Usetimes() func() {
	t := time.Now()
	pc, file, line, ok := runtime.Caller(1)
	_ = line
	_ = ok
	_ = file
	return func() {

		fmt.Printf("use time:%s, pc=%s, "+
			//"file=%s, " +
			//"line=%d, " +
			//"ok=%b" +
			"\n",
			time.Since(t).String(),
			runtime.FuncForPC(pc).Name(),
			//file,
			//line,
			//ok,
		)
	}
}

func GetDB() *gorm.DB {
	return db
}
