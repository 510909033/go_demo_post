package user

import "sync"

type user struct {
	Id       int
	Nickname string
}

var userPool = sync.Pool{
	New: func() interface{} {
		return &user{}
	},
}

func getUser() {
	userPool.Get()
}

/*
测试map情况下频繁更换user对象时的速度和gc情况
*/
func CreateSomeUser() {

}
