package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx := context.Background()

	ctx = context.WithValue(ctx, "basic_key", "basic_val")
	go dosome1(ctx)
	//go dosome1(ctx)
	//go dosome1(ctx)

	time.Sleep(time.Second * 10)
}

func dosome1(ctx context.Context) {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Printf("%s, ctx ptr = %p\n", runtime.FuncForPC(pc).Name(), ctx)

	ctx = context.WithValue(ctx, "key1", "val1_"+runtime.FuncForPC(pc).Name())
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	//_=cancelFunc
	go dosome11(ctx)

	select {
	case <-ctx.Done():
		fmt.Printf("dosome1 %v , done\n", ctx.Err())
	}
}

func dosome11(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Second*1)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("dosome11 %v done\n", ctx.Err())
		}
	}()
	time.Sleep(time.Second * 2)
	pc, _, _, _ := runtime.Caller(0)
	fmt.Printf("%s, ctx ptr = %p, val of key1=%s, basic_key=%s\n", runtime.FuncForPC(pc).Name(), ctx, ctx.Value("key1"), ctx.Value("basic_key"))
}
