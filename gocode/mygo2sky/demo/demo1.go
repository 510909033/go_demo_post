package demo

import (
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"go_demo_post/gocode/mygo2sky/config"
	"log"
	"time"
)

type d int

func (d d) ddd() d {
	log.Println("ddd")
	return d
}

func localspan(ctx context.Context) d {

	return d(1)
}

func aaa(ctx context.Context) {
	defer localspan(ctx).ddd()
	bbb(ctx)
}
func bbb(ctx context.Context) {
	func() {

	}()
}

func MySkywalkingDemo1() {
	//ctx := context.Background()
	//a := d(1)
	//_ = a
	//defer localspan(ctx).ddd()
	//
	//aaa(ctx)
	//aaa(ctx)

	log.Println("MySkywalkingDemo1")
	r, err := reporter.NewGRPCReporter(config.SERVER_ADDR)
	//r, err := reporter.NewLogReporter()
	//r,err:=reporter.NewLogReporter()
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer r.Close()
	tracer, err := go2sky.NewTracer("demo1", go2sky.WithReporter(r))

	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	// This for test
	span, ctx, err := tracer.CreateLocalSpan(context.Background())
	if err != nil {
		log.Fatalf("create new local span error %v \n", err)
	}

	//span.SetPeer("root_span_peer")
	//span.SetOperationName("MySkywalkingDemo1")
	span.Tag("tag_key", "tag_val")
	//span.Tag("span ptr", fmt.Sprintf("%p", ctx))
	fmt.Println("hha")
	_ = ctx
	time.Sleep(100 * time.Millisecond)
	//return

	reportedSpan := span.(go2sky.ReportedSpan)
	//log.Printf("reportedSpan=%+v", reportedSpan.StartTime(), reportedSpan.EndTime())

	span.End()
	dump(reportedSpan)
	time.Sleep(time.Second * 2)
}

func dump(s go2sky.ReportedSpan) {
	log.Println(s.StartTime(), s.EndTime(), s.Tags())
	log.Printf("%#v", s.Context())
	log.Printf("%#v", s.Context().FirstSpan.(go2sky.ReportedSpan))
}
