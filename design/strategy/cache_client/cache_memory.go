package cache_client

import (
	"go_demo_post/design/interf"
)

type CacheMemory struct {

}

func (c *CacheMemory) Get(key interf.IKey) (interf.IResult, error) {
	panic("implement me")
}

func (c *CacheMemory) Set(key interf.IKey, result interf.IResult) (bool, error) {
	panic("implement me")
}

