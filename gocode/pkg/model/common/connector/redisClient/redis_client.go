package redisClient

import (
	"github.com/go-redis/redis"
	"pkg/model/common/logger"
	"time"
)

func GetRedisClusterClient() *redis.ClusterClient {
	options := &redis.ClusterOptions{
		Addrs:              []string{"127.0.0.1:6379"},
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		OnNewNode:          nil,
		OnConnect:          nil,
		Password:           "",
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	client := redis.NewClusterClient(options)
	return client
}

func SetNx(key string, value interface{}, expiration time.Duration) {
	boolCmd := GetRedisClusterClient().SetNX(key, value, expiration)
	logger.GetLogger().Debugf("SetNx boolCmd=%+v\n", boolCmd)
}

func Set(key string, value interface{}, expiration time.Duration) {
	boolCmd := GetRedisClusterClient().Set(key, value, expiration)
	logger.GetLogger().Debugf("SetNx boolCmd=%+v\n", boolCmd)
}

func init() {
	//redis.Client{}.z
	Set("redis_s`et", "val", time.Second)

	//var c redis.Client
	//redis.StatusCmd{}
	//redis.IntCmd{}.
	//c.Set()
	//c.SetNX()
	//c.SetXX()
	//c.Expire()
	//c.
	//c.setr

}
