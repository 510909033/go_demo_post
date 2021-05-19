package strategy

import (
	"go_demo_post/design/interf"
)

type CacheService struct {

}

func (c *CacheService) Get(key interf.IKey) (interf.IResult, error) {
	panic("implement me")
}

func (c *CacheService) Set(key interf.IKey, result interf.IResult) (bool, error) {
	panic("implement me")
}

