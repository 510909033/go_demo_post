package main

import (
	"baotian0506.com/39_config/gocode/mygorm/demo_data"
)

func main() {
	//ctx:= context.Background()
	var demoData = &demo_data.DemoData{}
	_ = demoData
	demoData.InsertUser()
	//demoData.MultiInsertUser()
	//demoData.Walk(ctx)
	//mysqltest.Mysqltest1()
	//mysqltest.Mysqltest4()

	//	var baseCmd *redis.Cmd
	//var cmd *redis.StringCmd
	//redis.Client{}.Exists()
	//_=cmd
	//_=baseCmd
	//cmd.Err()
	//baseCmd.Result()
	//baseCmd.Err()
	//redis.BoolCmd{}

}
