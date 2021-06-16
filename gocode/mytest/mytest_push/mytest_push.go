package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("haha")
	t1("1", "2", "3")
}

func Test() {
	//	url = "http://go.wangbaotian.babytree-dev.com/go_preg_push/go_internal/push/send_push?push_config_id=1&process_id=0&target_id=12345&expected_send_ts=123&task_id=100&message=test+send+message&url=url%3A%2F%2Fsomeurl&extra=%7B%22frequency_config%22%3A%7B%22frequency_config_field%22%3A%7B%22group_id%22%3A%22group_12346%22%2C%22discussion_id%22%3A678%7D%7D%7D"
	//	http.Get(url)

}

func t1(a ...string) {
	fmt.Println(a)

	b, _ := json.Marshal(content)
	json.Unmarshal(b)
	//set cache
	//hset status 1进行中， 2已完成缓存设置
	//hset update_ts 操作的时间
	//save db

	//update cache ok

	//get from redis
	//分布式lock and get from db

	//time.Unix().Format()

}
