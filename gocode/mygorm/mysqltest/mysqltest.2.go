package mysqltest

import (
	"fmt"
	"gorm.io/gorm"
)

// UserInfo 用户信息
type UserInfo2 struct {
	gorm.Model
	Name   string
	Gender string
	Hobby  string
}

func Mysqltest2() {

	// 自动迁移
	db.AutoMigrate(&UserInfo2{})

	u1 := UserInfo2{gorm.Model{}, "七米", "男", "篮球"}

	// 创建记录
	db.Create(&u1)

	// 查询
	var u = new(UserInfo2)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo2
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
