package my_json

import (
	"encoding/json"
	"go_demo_post/common"
	"log"
	"time"
)

var str1 = `{"8":{"id":"8","title":"时尚"}}`

type Data1 map[string]DataOne
type DataOne struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Data2 map[string]DataOne2
type DataOne2 struct {
	Id    int    `json:"id"` //报错， 原始是string
	Title string `json:"title"`
}

type Data3 map[string]DataOne3
type DataOne3 struct {
	Id    json.Number `json:"id"`
	Title string      `json:"title"`
}

func MyJson() {

	go func() {
		defer common.ErrRecover()
		var data Data1
		err := json.Unmarshal([]byte(str1), &data)
		common.ErrPanic(err)
		log.Printf("%+v", data)
		//output
		// map[8:{Id:8 Title:时尚}]
	}()
	time.Sleep(time.Millisecond * 100)

	go func() {
		defer common.ErrRecover()
		var data Data2
		err := json.Unmarshal([]byte(str1), &data)
		common.ErrPanic(err)
		log.Printf("%+v", data)
		//output
		// cannot unmarshal string into Go struct field DataOne2.id of type int
	}()
	time.Sleep(time.Millisecond * 100)

	go func() {
		defer common.ErrRecover()
		var data Data3
		err := json.Unmarshal([]byte(str1), &data)
		common.ErrPanic(err)
		log.Printf("%+v", data)
		//output
		// map[8:{Id:8 Title:时尚}]
	}()
	time.Sleep(time.Millisecond * 100)

}
