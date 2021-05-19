package demo_rpc_json

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
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

type HelloService2 struct{}

func (p *HelloService2) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (p *HelloService2) Hello2(request map[string]string, reply *map[string]string) error {
	log.Println("request = ", request)
	args := make(map[string]string)
	args["version"] = "8.2.02"
	args["user_id"] = "2020"
	*reply = args
	//*reply = "hello:" + request
	time.Sleep(time.Second * 3) //
	return nil
}
func server() {

	rpc.RegisterName("HelloService", new(HelloService))
	rpc.RegisterName("HelloService2", new(HelloService2))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		//go rpc.ServeConn(conn)
	}
}

func client() {
	//client, err := rpc.Dial("tcp", "localhost:1234")
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//var reply string

	for {
		go func() {
			var reply map[string]string
			args := make(map[string]string)
			args["version"] = "8.1.0"
			args["user_id"] = "100"
			err := client.Call("HelloService.Hello2", args, &reply)

			if err != nil {
				log.Println("Client ERROR ", err)
			}
			log.Println("reply = ", reply)
		}()
		time.Sleep(time.Millisecond * 100)
	}
}

func DemoRpcJson() {

	go server()

	time.Sleep(time.Second * 1)

	//go client()

	select {}
}
