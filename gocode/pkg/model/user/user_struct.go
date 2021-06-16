package user

import (
	"gorm.io/gorm"
	"time"
)

type RegInfo struct {
	RegIp     string `gorm:"not null"`
	RegTimeTs int    `gorm:"autoCreateTime;not null"`
}

type UserInfo struct {
	ID        uint      `gorm:"primarykey;not null;autoIncrement"`
	CreatedAt time.Time `gorm:"not null"`
	// 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	UpdatedAt      int64 `gorm:"not null"`
	UpdatedNanoTs  int64 `gorm:"autoUpdateTime:nano;not null"`  // 使用时间戳填纳秒数充更新时间
	UpdatedMilliTs int64 `gorm:"autoUpdateTime:milli;not null"` // 使用时间戳毫秒数填充更新时间
	CreatedTs      int64 `gorm:"autoCreateTime;not null"`       // 使用时间戳秒数填充创建时间
	gorm.DeletedAt
	// Name 姓名
	Name string `gorm:"not null;size:100;uniqueindex"`
	// Gender 姓名
	Gender string `gorm:"not null;size:20;default:;"`
	// RegInfo 注册信息
	//
	// 嵌入结构体
	RegInfo          RegInfo `gorm:"embedded"`
	CountryCn        string  `gorm:"not null"`
	CountryEn        string  `gorm:"not null"`
	CountryPhoneCode string  `gorm:"not null"`
}
