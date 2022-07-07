package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"baotian0506.com/39_config/gocode/myredis/myredistest"
)

func main() {
	f, err := os.OpenFile("/tmp/redis_info.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	var wr []io.Writer
	wr = append(wr, os.Stdout)
	if err == nil {
		wr = append(wr, f)
	} else {
		panic(err)
	}

	log.SetOutput(io.MultiWriter(wr...))

	t := time.Now()
	//myredistest.SetMulti(100000)
	//myredistest.AddData(100)
	myredistest.TestRedisShutdown()
	//myredistest.GetId()

	fmt.Printf("time:%s\n", time.Since(t))

	//
	//	stringCmd:=client.Get(k)
	//
	//fmt.Printf("%#v, \n%#v\n",stringCmd,stringCmd.Err())
	//if stringCmd.Err() == redis.Nil{
	//	fmt.Println("==redis.nil")
	//}
	//
	//if errors.Is(stringCmd.Err(),redis.Nil) {
	//	fmt.Println("is redis.nil")
	//}

	//fmt.Println(client.Exists(k).Err())
	//client.ZRevRangeByScoreWithScores()
	//	c:=client.HGetAll(k)
	//
	//	fmt.Printf("%#v, \n%#v\n",c,c.Err())
}
