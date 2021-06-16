package user

import (
	"fmt"
	"hash/crc32"
	"pkg/model/common/connector/mysqlClient"
)

type UserSerevice struct {
}

/*
根据 手机号 登录， 【需要查询出 手机号 映射的库和表】

根据 ID 获取用户信息 【需要查询出 手机号 映射的库和表】
根据 ID 修改用户信息 【需要查询出 手机号 映射的库和表】


*/
/*
	//防止Name重复注册，
	1：唯一索引 ， 不适合分表分库
		1.1：解决办法 Name做crc32算法， 算出所在库和表 【不能修改数据库和表数量，否则映射会出错误】


*/
const (
	DbCount    = 2
	TableCount = 10
)

func GetTableNameByUserID(id uint64) string {
	return ""
}

//
func GetTableNameByName(name string) string {
	time33 := crc32.ChecksumIEEE([]byte(name))
	dbIndex := time33 % uint32(DbCount)
	tableIndex := time33 % uint32(TableCount)

	return fmt.Sprintf("%d_%d", dbIndex, tableIndex)
}
func GetTableNameByPhone(phone uint64) string {
	time33 := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%d", phone)))
	dbIndex := time33 % uint32(DbCount)
	tableIndex := time33 % uint32(TableCount)

	return fmt.Sprintf("%d_%d", dbIndex, tableIndex)
}

func (s UserSerevice) Insert() (*UserInfo, error) {
	gormdb, err := mysqlClient.GetDb()
	if err != nil {
		return nil, err
	}

	userInfo := &UserInfo{}

	userInfo.Name = "nickname"
	tx := gormdb.Create(userInfo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return userInfo, nil
}

func Demo() {
	//gdb:=mysqlClient.GetDb()

	//gdb.

}
