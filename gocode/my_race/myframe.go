package my_race

import (
	"context"
	"log"
	"runtime"
	"sort"
	"strings"
	"sync/atomic"
	"time"
)

type funcMap map[string]uint64
type MethodList struct {
	Name  string
	Count uint64
	Ts    time.Duration
}

type Ts struct {
}
type Monitor struct {
	hz         int
	StartTime  time.Time          `json:"-"`
	FuncMap    funcMap            `json:"-"` //函数名 调用次数
	Ticker     *time.Ticker       `json:"-"`
	Count      uint64             //执行了多少次 统计
	Cancel     context.CancelFunc `json:"-"` //go协程的取消函数
	Down       chan struct{}      `json:"-"`
	MethodList []*MethodList      //每个方法的耗时
	TotalTs    time.Duration      //总耗时
	SelectTs   time.Duration      //monitor.Ticker.C
	DownTs     time.Duration      //downFn函数的耗时

	EverySelectTs []int //每一个SelectTs的耗时
}

type methodListSlice []*MethodList

func (m methodListSlice) Len() int {
	return len(m)
}

func (m methodListSlice) Less(i, j int) bool {
	return m[i].Count < m[j].Count
}

func (m methodListSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func StartMonitor(ctx context.Context) *Monitor {
	//每个 frame.Function 出现的次数
	const hz = int(time.Millisecond) / 10
	//time.Second  1000 00 00 00
	EverySelectTsSliceLen := uint64(1000)
	monitor := &Monitor{
		hz:            hz,
		StartTime:     time.Now(),
		FuncMap:       make(map[string]uint64),
		Ticker:        time.NewTicker(time.Duration(hz)),
		Down:          make(chan struct{}, 0),
		EverySelectTs: make([]int, EverySelectTsSliceLen),
	}

	var newCtx context.Context
	newCtx, monitor.Cancel = context.WithCancel(ctx)

	downFn := func() {
		start := time.Now()
		defer func() {
			monitor.DownTs = time.Since(start)
		}()

		methodList := make(methodListSlice, len(monitor.FuncMap))
		index := 0
		for k, v := range monitor.FuncMap {
			methodList[index] = &MethodList{
				Name:  k,
				Count: v,
				Ts:    time.Duration(monitor.Count * uint64(monitor.hz)),
			}
			index++
		}
		sort.Sort(methodList)
		monitor.MethodList = methodList

		log.Println("before monitor.Down <- struct{}{}")
		monitor.Down <- struct{}{}
		log.Println("after monitor.Down <- struct{}{}")
	}

	go func() {
		startTsGo := time.Now()
		defer func() {
			monitor.TotalTs = time.Since(startTsGo)
		}()
		for {
			select {
			case <-newCtx.Done():
				log.Println("newCtx.Done")
				downFn()
				return
			case <-ctx.Done():
				log.Println("ctx2.Done")
				downFn()
				return
			case <-monitor.Ticker.C:
				//start := time.Now()
				//atomic.AddUint64(&monitor.Count, 1)
				//continue
				start := time.Now()
				atomic.AddUint64(&monitor.Count, 1)
				rpc := make([]uintptr, 1)
				n := runtime.Callers(1, rpc)
				if n < 1 {
					panic("could not determine selfFrame")
				}
				selfFrame, _ := runtime.CallersFrames(rpc).Next()

				var p []runtime.StackRecord
				n, ok := runtime.GoroutineProfile(nil)
				for {
					// Allocate room for a slightly bigger profile,
					// in case a few more entries have been added
					// since the call to ThreadProfile.
					p = make([]runtime.StackRecord, n+10)
					n, ok = runtime.GoroutineProfile(p)
					//log.Println("len(p)=%d, n=%d, ok=%b", len(p), n, ok)
					if ok {
						p = p[0:n]
						break
					}
					// Profile grew; try again.
				}

			outer:
				for _, pp := range p {
					frames := runtime.CallersFrames(pp.Stack())

					var stack []string
					for {
						frame, more := frames.Next()
						if !more {
							break
						} else if frame.Entry == selfFrame.Entry {
							continue outer
						}

						stack = append([]string{frame.Function}, stack...)
					}
					key := strings.Join(stack, ";")
					monitor.FuncMap[key]++
				}
				monitor.SelectTs += time.Since(start)
				monitor.EverySelectTs[monitor.Count%EverySelectTsSliceLen] = int(time.Since(start))
			} //end for

			//time.Sleep(time.Millisecond * 5)

			//os.Exit(1)
		} //end for
	}()

	return monitor
}

func myFrame2() {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(1, rpc)
	if n < 1 {
		panic("could not determine selfFrame")
	}
	selfFrame, _ := runtime.CallersFrames(rpc).Next()

	var p []runtime.StackRecord
	n, ok := runtime.GoroutineProfile(nil)
	for {
		// Allocate room for a slightly bigger profile,
		// in case a few more entries have been added
		// since the call to ThreadProfile.
		p = make([]runtime.StackRecord, n+10)
		n, ok = runtime.GoroutineProfile(p)
		log.Println("len(p)=%d, n=%d, ok=%b", len(p), n, ok)
		if ok {
			p = p[0:n]
			break
		}
		// Profile grew; try again.
	}

outer:
	for _, pp := range p {
		frames := runtime.CallersFrames(pp.Stack())

		var stack []string
		for {
			frame, more := frames.Next()
			if !more {
				break
			} else if frame.Entry == selfFrame.Entry {
				continue outer
			}

			stack = append([]string{frame.Function}, stack...)
		}
		key := strings.Join(stack, ";")
		_ = key
		//s[key]++
	}
}

/*
//rpc := make([]uintptr, 100)
				//n1 := runtime.Callers(0, rpc)
				//_ = n1
				////log.Println("n1=", n1)
				//frames := runtime.CallersFrames(rpc) //.Next()
				//log.Println(frames)

				//for {
				//	if frame, more := frames.Next(); more {
				//		monitor.FuncMap[frame.Function]++
				//log.Printf("frame=%#v", frame)
				//log.Printf("frame.Line = %d ,frame.Function=%s, frame.File = %s", frame.Line, frame.Function, frame.File)
				//log.Printf("frame.Func.Name()=%s", frame.Func.Name())
				//file, line := frame.Func.FileLine(frame.PC)
				//log.Printf("frame.Func.FileLine(frame.PC), file=%s , line=%d", file, line)
				//
				//file, line = frame.Func.FileLine(frame.Entry)
				//log.Printf("frame.Func.FileLine(frame.Entry), file=%s , line=%d", file, line)
				//log.Println("")
				//	} else {
				//		break
				//	}
				//} //end for
*/
