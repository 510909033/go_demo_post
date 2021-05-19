package myoption

import (
	"context"
	"time"
)

type Option struct {
	num       int
	timeout   time.Duration
	openCache bool
	debug     bool
}

type IOption interface {
	apply(option *Option)
}

type DemoOption struct {
}

type fn func(option *Option)

func (f fn) apply(option *Option) {
	f(option)
}

func WithTimeout(d time.Duration) IOption {
	return fn(func(option *Option) {
		option.timeout = d
	})
}

func NewDemoOption(ctx context.Context, opt ...IOption) *DemoOption {
	option := &Option{
		num:       0,
		timeout:   0,
		openCache: false,
		debug:     false,
	}

	for _, o := range opt {
		o.apply(option)
	}

	return nil
}
