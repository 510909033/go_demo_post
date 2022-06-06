package mygo2sky

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
)

//func main() {
//	example4()
//}

//func example4() {
//	log.SetFlags(log.Lshortfile)
//	ctx := context.Background()
//
//	client.Start(ctx)
//
//	select {}
//
//}

func example2() {
	// Use gRPC reporter for production
	//r, err := reporter.NewGRPCReporter("172.16.10.43:11800")
	r, err := reporter.NewGRPCReporter("192.168.6.3:11800")
	//r,err:=reporter.NewLogReporter()
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer r.Close()
	tracer, err := go2sky.NewTracer("example", go2sky.WithReporter(r))

	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	// This for test
	span, ctx, err := tracer.CreateLocalSpan(context.Background())
	if err != nil {
		log.Fatalf("create new local span error %v \n", err)
	}

	span.SetPeer("root_span_peer")
	span.SetOperationName("/debug/api/demo")
	span.Tag("tag_key", "tag_val")
	span.Tag("span ptr", fmt.Sprintf("%p", ctx))
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		stri := fmt.Sprintf("%d", i)
		subSpan, ctx1, err := tracer.CreateLocalSpan(ctx)
		if err != nil {
			log.Fatalf("create new sub local span error %v \n", err)
		}
		subSpan.SetPeer("sub_span_peer")
		subSpan.Tag(go2sky.Tag(fmt.Sprintf("tag_%d", i)), fmt.Sprintf("val_%d", i))
		subSpan.Tag(go2sky.Tag("span ptr_"+stri), fmt.Sprintf("%p", ctx1))
		subSpan.SetOperationName("operation_name_" + stri)
		subSpan.Log(time.Now(), "我是log"+stri)
		time.Sleep(time.Millisecond * 100 * (time.Duration(i)))

		var wg sync.WaitGroup

		wg.Add(1)
		go func(ctx context.Context, i int) {
			stri := fmt.Sprintf("%d", i)
			subSpan, ctx1, _ := tracer.CreateLocalSpan(ctx)
			defer subSpan.End()
			defer wg.Done()
			subSpan.SetPeer("sub_sub_span_peer")
			subSpan.Tag(go2sky.Tag(fmt.Sprintf("sub_tag_%d", i)), fmt.Sprintf("sub_val_%d", i))
			subSpan.Tag(go2sky.Tag("sub_span ptr_"+stri), fmt.Sprintf("%p", ctx1))
			subSpan.SetOperationName("sub_operation_name_" + stri)
			subSpan.Log(time.Now(), "sub_我是log"+stri)
			time.Sleep(time.Millisecond * 100 * (time.Duration(i)))

		}(ctx1, i)
		wg.Wait()
		subSpan.End()
	}

	time.Sleep(500 * time.Millisecond)
	span.End()

	time.Sleep(time.Second * 2)
	// Output:
	fmt.Println("over")
}

var tracer *go2sky.Tracer
var wg sync.WaitGroup

func init() {
	log.Println("ExampleNewTracer 需要取消下面的return")
	return
	// Use gRPC reporter for production

	//r, err := reporter.NewGRPCReporter("172.16.10.43:11800")
	r, err := reporter.NewGRPCReporter("172.16.7.186:11800")
	//r1, err := reporter.NewLogReporter()

	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	//defer r.Close()
	tracer, err = go2sky.NewTracer("example", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
}

//18
func ExampleNewTracer() {
	// This for test
	span, ctx, err := tracer.CreateLocalSpan(context.Background())
	if err != nil {
		log.Fatalf("create new local span error %v \n", err)
	}

	span.SetOperationName("ExampleNewTracer")
	span.Tag("kind", "root")

	wg.Add(1)
	go level1(ctx)

	wg.Add(1)
	go level2(ctx)

	wg.Wait()
	span.End()
	fmt.Println("over")
	time.Sleep(time.Second * 3)
}

func level1(ctx context.Context) {

	span, ctx, err := tracer.CreateLocalSpan(ctx)
	if err != nil {
		log.Fatalf("create new local span error %v \n", err)
	}
	defer wg.Done()
	defer span.End()

	span.SetOperationName("level_1")
	time.Sleep(time.Millisecond * 200)
}

func level2(ctx context.Context) {
	span, ctx, err := tracer.CreateLocalSpan(ctx)
	if err != nil {
		log.Fatalf("create new local span error %v \n", err)
	}

	defer wg.Done()
	defer span.End()

	span.SetOperationName("level_2")
	time.Sleep(time.Millisecond * 400)
}

func example3() {
	// Use gRPC reporter for production

	//r, err := reporter.NewGRPCReporter("172.16.10.43:11800")
	r, err := reporter.NewGRPCReporter("192.168.6.3:11800")

	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer r.Close()
	//tracer, err := go2sky.NewTracer("example", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	//
	//
	//span,ctx,err:=tracer.CreateEntrySpan(context.Background(),"operation_name,",func() (string,error){
	//	return "",nil
	//})
	//if err != nil {
	//	log.Fatalf("CreateEntrySpan error %v \n", err)
	//}
	//time.Sleep(time.Millisecond*100)

	//// This for test
	//span1, err := tracer.CreateExitSpan(context.Background(), "operation_name1", "peer_name1",
	//	func(header string) error {
	//		return nil
	//	})
	//if err != nil {
	//	log.Fatalf("CreateEntrySpan error %v \n", err)
	//}
	//time.Sleep(time.Millisecond * 200)
	//span1.End()
	//time.Sleep(time.Millisecond*300)
	//span.End()

	//span1.SetOperationName("invoke data")
	//span.Tag("kind", "outer")
	//time.Sleep(time.Second)
	//span.End()
	//return

	//ctx:=context.Background()
	//
	//
	//
	//subSpan, _, err := tracer.CreateLocalSpan(ctx)
	//if err != nil {
	//	log.Fatalf("create new sub local span error %v \n", err)
	//}
	//subSpan.SetOperationName("invoke inner")
	//subSpan.Log(time.Now(), "inner", "this is right")
	//time.Sleep(time.Second)
	//subSpan.End()
	//time.Sleep(500 * time.Millisecond)
	//span.End()
	//time.Sleep(time.Second)
	//// Output:
	fmt.Println("over")
}
