package mysqltest

import (
	"baotian0506.com/39_config/gocode/mygorm/mycommon"
	"gorm.io/gorm"
)

var db *gorm.DB = mycommon.GetDB()
var usetimes func() func() = mycommon.Usetimes
