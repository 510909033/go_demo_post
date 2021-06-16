package mysqltest

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// UserInfo 用户信息
type UserInfo4 struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt
	Name     string
	Gender   string
	Hobby    string
	Phone    string
	AInt     uint64
	BInt     uint32
	CInt     int32
	DInt     uint8
	EInt     int8
	FByte    byte
	GFloat   float32
	HFloag64 float64
	ANotNull string `gorm:"not null,size:1000,unique_index"`
	BNotNull string `gorm:"not null;size:1000;unique_index"`
}

func Mysqltest4() {
	//multiInsertSync()
	//multiInsertAsync()
	//return

	defer usetimes()()
	// 自动迁移
	db.AutoMigrate(&UserInfo4{})

	u1 := UserInfo4{Name: "七米", Gender: "男", Hobby: "篮球"}

	// 创建记录
	db.Create(&u1)

	// 查询
	var u = new(UserInfo4)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo4
	db.Limit(1).Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}

func multiInsertAsync() {
	defer usetimes()()

	ch := make(chan bool, 50)
	var i int32 = 10000
	var wg sync.WaitGroup
	myrand := &MyRand{}

	for {
		atomic.AddInt32(&i, -1)
		if i < 0 {
			break
		}
		wg.Add(1)
		ch <- true
		go func() {
			defer wg.Done()
			defer func() {
				<-ch
				//fmt.Println("???")
			}()

			u := UserInfo4{}
			u.Name = myrand.GetCn(10)
			u.Gender = "男"
			if i%5 == 0 {
				u.Gender = "女"
			}
			rand.Seed(time.Now().UnixNano())
			u.Phone = fmt.Sprintf("%s", rand.Intn(100000000))
			db.Create(&u)
			//time.Sleep(time.Second)
		}()
	}

	wg.Wait()
	fmt.Println("over")
}

func multiInsertSync() {
	var i int32 = 10

	for {

		atomic.AddInt32(&i, -1)
		if i < 1 {
			break
		}

		u := UserInfo4{}
		tx := db.Create(&u)
		_ = tx

	}

}
