package mysqlClient

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetDb() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}
	//db.Debug()
	//defer db.Close()

	sqlDB, _ := db.DB()
	//sqlDB.Close()
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetConnMaxLifetime(time.Second)

	go func() {
		for {
			//todo
			//fmt.Printf("%+v\n",sqlDB.Stats())
			time.Sleep(time.Second)
		}
	}()
	return db, err
}
