package myredistest

import "testing"

func TestGetId(t *testing.T) {
	for i := 0; i < 1; i++ {
		GetId()
	}
}

func TestMset(t *testing.T) {
	Mset()
}

func TestLuaSomeMethod(t *testing.T) {
	LuaSomeMethod()
}

func TestLuaSomeMethod2(t *testing.T) {
	LuaSomeMethod2()
}

func TestMget(t *testing.T) {
	Mget()
}

func TestLuaDel(t *testing.T) {
	LuaDel()
}