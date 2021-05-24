package my_json

import (
	"bufio"
	"bytes"
	"encoding/json"
	"time"
)

type User struct {
	Id   int64
	Name string
	Desc string
	Tags []string
}

var user = getUser()
var buf = make([]byte, 1024)

func DemoMyJson1() string {

	buffer := bytes.NewBuffer(buf)
	writer := bufio.NewWriter(buffer)

	writer.WriteString("")
	writer.WriteString("")
	writer.WriteString("")
	writer.Flush()
	return buffer.String()
}

func DemoMyJson2() string {
	s, _ := json.Marshal(user)
	return string(s)
}
func getUser() *User {
	return &User{
		Id:   time.Now().Unix(),
		Name: "aaaaaaaaaaaa",
		Desc: "bbbbbbbbbbb",
		Tags: []string{"a", "b", "c"},
	}
}
