package mysqltest

import (
	"fmt"
	"time"
)

// UserInfo 用户信息
type UserInfo3 struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Gender    string
	Hobby     string
}

func Mysqltest3() {

	// 自动迁移
	db.AutoMigrate(&UserInfo3{})

	u1 := UserInfo3{Name: "七米", Gender: "男", Hobby: "篮球"}

	// 创建记录
	db.Create(&u1)

	// 查询
	var u = new(UserInfo3)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo3
	db.Limit(1).Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
