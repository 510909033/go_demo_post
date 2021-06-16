package mysky

import (
	"baotian0506.com/39_config/gocode/mygin/common"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/propagation"
	"github.com/SkyAPM/go2sky/reporter"
	language_agent "github.com/SkyAPM/go2sky/reporter/grpc/language-agent"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

var tracer = GetTracer("mygo2skyclient")

const (
	SERVER_ADDR                       = "192.168.6.3:11800"
	httpClientComponentID       int32 = 2
	mysqlComponentID            int32 = 5
	redisComponentID            int32 = 7
	memcachedComponentID        int32 = 20
	httpServerComponentID       int32 = 49
	HttpServerComponentID       int32 = 49
	rpcComponentID              int32 = 50
	rabbitmqComponentID         int32 = 51
	rabbitmqProducerComponentID int32 = 52
	rabbitmqConsumerComponentID int32 = 53
)

var traceMap = make(map[string]*go2sky.Tracer)
var once = sync.Once{}
var mu sync.Mutex

func getMacAddrs() string {
	macAddrs := make([]string, 0)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return ""
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}

	hash := md5.New()
	//hash.Write([]byte(strings.Join(macAddrs,"|")))
	return fmt.Sprintf("%x", hash.Sum([]byte(strings.Join(macAddrs, "|"))))
}

func GetTracer(serviceName string) *go2sky.Tracer {

	mu.Lock()
	defer mu.Unlock()
	if _, ok := traceMap[serviceName]; ok {
		return traceMap[serviceName]
	}

	//if r == nil{
	r1, err := reporter.NewGRPCReporter(SERVER_ADDR)
	if err != nil {
		log.Fatalf("[New GRPC Reporter Error]: [%v]", err)
		common.Exception(err)
		return nil
	}
	//}
	//tracer, err = go2sky.NewTracer("client_test", go2sky.WithReporter(r), go2sky.WithInstance("RTS_Test_1"))
	tracer1, err := go2sky.NewTracer(serviceName, go2sky.WithReporter(r1), go2sky.WithInstance(serviceName+"_"+getMacAddrs()))
	if err != nil {
		log.Fatalf("[New Tracer Error]: [%v]", err)
		common.Exception(err)
		return nil
	}
	traceMap[serviceName] = tracer1

	return traceMap[serviceName]
}
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("init")
}

func commonSpan(ctx context.Context) (context.Context, go2sky.Span, func()) {
	pc, file, line, ok := runtime.Caller(1)
	_, _, _ = file, line, ok

	var span go2sky.Span
	var err error
	span, ctx, err = tracer.CreateLocalSpan(ctx)

	if err != nil {
		common.Exception(err)
	}
	span.SetPeer("peer_" + runtime.FuncForPC(pc).Name())
	span.SetOperationName("opt_" + runtime.FuncForPC(pc).Name())
	//span.Tag("tes", runtime.FuncForPC(pc).)
	return ctx, span, func() {
		span.End()
	}
}

func Start() {
	var wg sync.WaitGroup
	ch := make(chan int, 200)
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		ch <- 1
		go func() {
			defer func() {
				<-ch
				wg.Done()
			}()
			ctx := context.Background()
			ctx, span, fn := commonSpan(ctx)
			_ = span
			//span.SetPeer("my_client")

			log.Println("start")

			v1Login(ctx)
			v2Login(ctx)
			v1Submit(ctx)

			fn()
		}()
	}

	wg.Wait()
	time.Sleep(time.Second * 2)
}

func v1Login(ctx context.Context) {
	ctx, _, fn := commonSpan(ctx)
	defer fn()
	client := &http.Client{}
	client.Timeout = time.Second * 5
	url := "http://127.0.0.1:11111/v1/login"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	var r = reqest

	span, err := tracer.CreateExitSpan(ctx, r.Method+"/"+reqest.URL.Path, reqest.Host, func(header string) error {
		reqest.Header.Add(propagation.Header, header)
		return nil
	})
	common.Exception(err)
	span.SetSpanLayer(language_agent.SpanLayer_Http)
	span.SetComponent(httpClientComponentID)
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.Tag(go2sky.TagURL, fmt.Sprintf("%s%s", r.Host, r.URL.Path))

	defer span.End()

	//增加header选项

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()
}

func v1Submit(ctx context.Context) {
	client := &http.Client{}
	client.Timeout = time.Second * 5
	url := "http://127.0.0.1:11111/v1/submit"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	var r = reqest

	span, err := tracer.CreateExitSpan(ctx, r.Method+"/"+reqest.URL.Path, reqest.Host, func(header string) error {
		reqest.Header.Add(propagation.Header, header)
		return nil
	})
	common.Exception(err)
	span.SetSpanLayer(language_agent.SpanLayer_Http)
	span.SetComponent(httpClientComponentID)
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.Tag(go2sky.TagURL, fmt.Sprintf("%s%s", r.Host, r.URL.Path))

	defer span.End()

	//增加header选项

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	common.Exception(err)
	defer response.Body.Close()
}

const PEER_V1 = "192.168.0.5"

func v2Login(ctx context.Context) {
	client := &http.Client{}
	url := "http://127.0.0.1:22222/v2/login"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	var r = reqest
	span, err := tracer.CreateExitSpan(ctx, r.Method+"/"+reqest.URL.Path, reqest.Host, func(header string) error {
		reqest.Header.Add(propagation.Header, header)
		return nil
	})
	common.Exception(err)

	span.SetSpanLayer(language_agent.SpanLayer_Http)
	span.SetComponent(httpClientComponentID)
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.Tag(go2sky.TagURL, fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	//span.Tag(go2sky.TagStatusCode, strconv.Itoa(200))//不加这个，加error
	//span.Error(time.Now(), err.Error()) //用这个
	common.Exception(err)
	defer span.End()

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()
}
