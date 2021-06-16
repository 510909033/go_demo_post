// demo_data包
//
// 提供公共的数据生成方法
//
// 写入数据到mysql数据库
//
// 生成nginx 日志，
// sh genLog.sh
//
package demo_data

import (
	"baotian0506.com/39_config/gocode/mygorm/bgdata"
	"context"
)

var user = &bgdata.UserInfo{}
var fileRange = bgdata.NewFileRange()

type DemoData struct {
}

func (s *DemoData) InsertUser() {
	user.MyInsertOne()
	user.MyInsertOne()
	user.MyInsertOne()
}

func (s *DemoData) MultiInsertUser() {
	for i := 0; i < 10000; i++ {
		user.MyInsertOne()
	}
}

func (s *DemoData) Walk(ctx context.Context) {
	root := `C:\Users\Administrator\go\src\baotian0506.com\39_config\gocode\mygorm\`
	fileRange.Walk(root)
}
