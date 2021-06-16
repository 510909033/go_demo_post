package main

import (
	"pkg/model/common/connector/redisClient"
	"pkg/model/common/unique_hashids"
	"time"
)

func main() {
	//go api.Register()
	unique_hashids.Demo()
	redisClient.Set("redis_set", "val", time.Second)
}
