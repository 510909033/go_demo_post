package demo_rpc_http

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"time"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	time.Sleep(time.Second * 3)
	*reply = "hello:" + request
	log.Println("hello.reply=", *reply)

	return nil
}

func (p *HelloService) Hello2(request map[string]string, reply *map[string]string) error {
	log.Println("request = ", request)
	args := make(map[string]string)
	args["version"] = "8.2.0"
	args["user_id"] = "200"
	*reply = args
	//*reply = "hello:" + request
	time.Sleep(time.Second * 3) //
	return nil
}

type MyRdCloser struct {
	io.Reader
}

func (MyRdCloser) Close() error { return nil }
func server() {

	//http://localhost:1234/jsonrpc?heihei={"method":"HelloService.Hello","params":["hello123311"],"id":110}

	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {

		val := r.URL.Query().Get("heihei")
		rdCloser := MyRdCloser{
			bufio.NewReader(strings.NewReader(val)),
		}
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: rdCloser,
			Writer:     w,
		}
		w.Header().Set("x-haha", "x-value")

		log.Println("val=", val)

		//rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}

func client() {

	//http.Get("http://")

}

func DemoRpcHttp() {

	go server()

	//time.Sleep(time.Second * 1)

	//go client()

	select {}
}
