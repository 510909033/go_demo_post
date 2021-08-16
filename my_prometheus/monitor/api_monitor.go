package monitor

import (
	"context"
	"time"
)

/*
监控api p99
同级的
子集的

api_monitor{"api"="/api/user/get_user_info"}

api_monitor{"api"="/api/user/get_user_info", level=1, level1="get_follow_info"}

api_monitor{"api"="/api/user/get_user_info", level=2, level1="get_follow_info", level2="/rpc/java/get_user_info"}


*/

const (
	END_FN_KEY = "end_fn"
)

type ApiMonitor interface {
	Sibling(currName string) ApiMonitor
	Sub(currName string) ApiMonitor

	End()
	GetRootName() string
	GetCurrName() string
	GetLevel() int
	GetConsuming() float64 //毫秒
}

func NewMonitor(ctx context.Context, root string, fn func(ApiMonitor)) ApiMonitor {
	return builder(ctx, root, root, 0, time.Now(), make(map[string]string))
}

func builder(ctx context.Context, root string, currName string, level int, start time.Time, lables map[string]string) ApiMonitor {
	if lables == nil {
		lables = make(map[string]string)
	}
	return &DemoMonitor{
		Ctx:      ctx,
		RootName: root,
		CurrName: currName,
		Level:    level,
		Start:    start,
		Labels:   lables,
	}
}

type DemoMonitor struct {
	Ctx      context.Context
	RootName string
	CurrName string
	Level    int
	Start    time.Time
	Labels   map[string]string
}

func (d *DemoMonitor) GetConsuming() float64 {
	return float64(time.Since(d.Start).Nanoseconds())
}

func (d *DemoMonitor) GetRootName() string {
	return d.RootName
}

func (d *DemoMonitor) GetCurrName() string {
	return d.CurrName
}

func (d *DemoMonitor) GetLevel() int {
	return d.Level
}

func (d *DemoMonitor) Sibling(currName string) ApiMonitor {
	return builder(d.Ctx, d.RootName, currName, d.Level, time.Now(), nil)
}

func (d *DemoMonitor) Sub(currName string) ApiMonitor {
	ctx := context.WithValue(d.Ctx, "currName", currName)
	return builder(ctx, d.RootName, currName, d.Level+1, time.Now(), nil)
}

func (d *DemoMonitor) End() {
	fn := d.Ctx.Value(END_FN_KEY)
	if fn, ok := fn.(func(monitor ApiMonitor)); ok {
		fn(d)
	}
}
