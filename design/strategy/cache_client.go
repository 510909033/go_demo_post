package strategy

import (
	"go_demo_post/design/interf"
)

type CacheClient struct {
	client interf.ICacheClient
}



func NewClient(client interf.ICacheClient) *CacheClient {
	return &CacheClient{client:client}
}

func (c *CacheClient) Get(key string) {
	c.client.Get(key)
}

func (c *CacheClient) Set(key, val string){
	panic("implement me")
}

