package client

import (
	"baotian0506.com/app/mygo2sky/config"
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	language_agent "github.com/SkyAPM/go2sky/reporter/grpc/language-agent"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Req struct {
	A      int
	Header string // 添加此字段，用于传递context信息
}

func (p *Req) Set(key, value string) error {
	p.Header = fmt.Sprintf("%s:%s", key, value)
	return nil
}

var H = ""
var r go2sky.Reporter
var err error
var tracer *go2sky.Tracer
var wg sync.WaitGroup

func init() {
	r, err = reporter.NewGRPCReporter(config.SERVER_ADDR)
	if err != nil {
		log.Fatalf("[New GRPC Reporter Error]: [%v]", err)
		return
	}

	//tracer, err = go2sky.NewTracer("client_test", go2sky.WithReporter(r), go2sky.WithInstance("RTS_Test_1"))
	tracer, err = go2sky.NewTracer("client_test", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("[New Tracer Error]: [%v]", err)
		return
	}
}

func GetTracer() *go2sky.Tracer {
	return tracer
}

func Start(ctx context.Context) {
	ctx, fn := config.GetLocalspan(ctx, tracer)
	defer fn(nil)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		OneClient(ctx)
		time.Sleep(getSleepTimeN(50))
	}

	wg.Wait()

	//time.Sleep(time.Second*3)
	//	//r.Close()
}
func getSleepTimeN(t int) time.Duration {
	return time.Duration(rand.Intn(t))
}
func getSleepTime() time.Duration {
	return time.Duration(rand.Intn(100) * rand.Intn(10))
}

//请求用户中心的登录服务
func login(ctx context.Context) {
	var h string
	var t = time.Now()
	span, err := tracer.CreateExitSpan(ctx, "/service/api/login", "peer_name1",
		func(header string) error {
			//todo set header
			h = header
			log.Printf("CreateExitSpan header=%s\n", header)
			return nil
		})
	if err != nil {
		log.Fatalf("CreateEntrySpan error %v \n", err)
	}
	log.Printf("login 创建CreateExitSpan消耗时间：%s\n", time.Since(t))

	t = time.Now()
	defer func() {
		log.Printf("login span.End消耗时间：%s\n", time.Since(t))
	}()
	defer span.End()

	//mock /service/api/login
	s := &Server{
		H: h,
	}
	s.Login()
}

func GetUserAddress(ctx context.Context) context.Context {
	span2, ctx2, _ := tracer.CreateLocalSpan(ctx)
	span2.SetPeer("span2_peer")
	span2.SetOperationName("GetUserAddress")

	time.Sleep(getSleepTime())
	span2.End()
	return ctx2
}
func OneClient(ctx context.Context) {
	defer wg.Done()

	ctx, fn := config.GetLocalspan(ctx, tracer)
	defer fn(nil)

	t := time.Now()

	span1, ctx, _ := tracer.CreateEntrySpan(ctx, "OnClient", func() (string, error) {
		return "", nil
	})
	span1.SetComponent(go2sky.ComponentIDHttpServer)
	span1.Tag(go2sky.TagHTTPMethod, "GET")
	span1.Tag(go2sky.TagStatusCode, "200")
	span1.Tag(go2sky.TagURL, "/h5/user/login")
	span1.SetPeer("")
	span1.SetSpanLayer(language_agent.SpanLayer_Http)

	log.Printf("OneClient创建CreateEntrySpan消耗时间 time=%s\n", time.Since(t))

	//登录
	log.Printf("login before time=%s\n", time.Since(t))
	login(ctx)
	log.Printf("login after time=%s\n", time.Since(t))

	//登录成功后获取用户地址
	log.Printf("GetUserAddress before time=%s\n", time.Since(t))
	t1 := time.Now()
	ctx2 := GetUserAddress(ctx)
	log.Printf("GetUserAddress 消耗时间 time=%s\n", time.Since(t1))
	_ = ctx2

	//span1.Error(time.Now(),"test_arror")

	defer func() {
		fmt.Printf("oneClient after time=%s\n", time.Since(t))
	}()
	span1.End()

}

type Server struct {
	H string
}

func (s *Server) Login() {
	r, err := reporter.NewGRPCReporter(config.SERVER_ADDR)
	if err != nil {
		log.Fatalf("[New GRPC Reporter Error]: [%v]", err)
		return
	}

	//tracer, err := go2sky.NewTracer("server_test", go2sky.WithReporter(r), go2sky.WithInstance("RTS_Test_1"))
	tracer, err := go2sky.NewTracer("server_test", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("[New Tracer Error]: [%v]", err)
		return
	}

	span1, ctx, err := tracer.CreateEntrySpan(context.Background(), "/Server/Login", func() (string, error) {
		fmt.Printf("server H=%s\n", s.H)
		return s.H, nil
	})
	if err != nil {
		log.Fatalf("[CreateEntrySpan Error]: [%v]", err)
		return
	}

	_ = ctx
	span1.SetPeer("server_peer")

	time.Sleep(getSleepTime())
	span1.End()
}
