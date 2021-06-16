package bgdata

import (
	"baotian0506.com/39_config/gocode/mygorm/mycommon"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	CN_MIN_UNICODE = 0x4e00
	CN_MAX_UNICODE = 0x9fa5
)

var bgdata_log_dir = "e:/logs"

// db mysql数据库
var db = mycommon.GetDB()

//myrand 提供了随机方法的对象
var myrand = &MyRand{}

//go的log包提供的记录日志功能的Logger
var logger *log.Logger

func init() {
	out, err := os.OpenFile(bgdata_log_dir+"/bgdata.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	multiOut := io.MultiWriter(out, os.Stdout)
	logger = log.New(multiOut, "", log.LstdFlags|log.Lshortfile)
}

// create 将一个struct对象通过gorm的create方法写入表中
//
//错误会打印日志
func create(val interface{}) {
	defer mycommon.Usetimes()()
	tx := db.Create(val)
	if tx.Error != nil {
		fmt.Printf("create err=%#v, val=%#v\n", tx.Error, val)
	}

}
