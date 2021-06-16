package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync/atomic"
	"time"
)

type Args struct {
	A, B int
}
type Reply struct {
	Data string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	t := time.Now()
	var cnt uint64 = 0
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("cnt= ", cnt, " ", time.Since(t).String())

		}
	}()
	for {
		args := &Args{7, 8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		//		fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
		atomic.AddUint64(&cnt, 1)
	}
}
