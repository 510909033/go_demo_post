package myredistest

import "github.com/go-redis/redis"

const ADDR = "192.168.6.3:6379"

var client *redis.Client

func init() {
	var opt = &redis.Options{
		Network:            "tcp",
		Addr:               ADDR,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
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
	client = redis.NewClient(opt)
}
