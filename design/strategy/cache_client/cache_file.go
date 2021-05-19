package cache_client

import (
	"go_demo_post/design/interf"
	"log"
)

type CacheFile struct {
}

func (c *CacheFile) Get(key interf.IKey) (interf.IResult, error) {
	log.Println(c)
	return nil,nil
}

func (c *CacheFile) Set(key interf.IKey, result interf.IResult) (bool, error) {
	panic("implement me")
}

