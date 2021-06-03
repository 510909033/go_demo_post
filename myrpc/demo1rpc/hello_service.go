package demo1rpc

import (
	"log"
	"net"
	"net/rpc"
	"sync"
	"sync/atomic"
	"time"
)

type HelloService struct {
	mu sync.Mutex
	Pv uint64
}

func (p *HelloService) monitor() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println(p.Pv)
			}
		}
	}()
}
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

var bigStr = [10000]byte{}

func (p *HelloService) Hello2(request map[string]string, reply *map[string]string) error {
	p.mu.Lock()
	atomic.AddUint64(&p.Pv, 1)
	p.mu.Unlock()
	//log.Println("request = ", request)
	args := make(map[string]string)
	args["version"] = "8.2.0"
	args["user_id"] = "200"
	args["big_data"] = string(bigStr[:10000])
	*reply = args
	//*reply = "hello:" + request
	//time.Sleep(time.Second * 3) //
	time.Sleep(time.Millisecond * 5)
	return nil
}

func server() {
	helloService := new(HelloService)
	helloService.monitor()
	rpc.RegisterName("HelloService", helloService)

	listener, err := net.Listen("tcp", "mylocalhost:1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}

func clientCall() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//var reply string
	var reply map[string]string
	for {
		args := make(map[string]string)
		args["version"] = "8.1.0"
		args["user_id"] = "100"
		err = client.Call("HelloService.Hello2", args, &reply)
		if err != nil {
			log.Println("Client ERROR ", err)
		}
		//log.Println("reply = ", reply)
		//time.Sleep(time.Millisecond * 10)
	}
}

func clientGo() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//var reply string
	var done = make(chan *rpc.Call, 2) //限流

	go func() {
		for {
			select {
			case call := <-done:
				log.Println("call.err", call.Error)
				//log.Println("call.reply", call.Reply)
			}
			//break
		}
	}()

	for {
		log.Println(len(done), cap(done))

		go func() {
			var reply map[string]string
			args := make(map[string]string)
			args["version"] = "8.1.0"
			args["user_id"] = "100"
			//call := client.Go("HelloService.Hello2", args, &reply, done)
			client.Go("HelloService.Hello2", args, &reply, done)

			//calls := <-call.Done
			//err := calls.Error
			//
			////err = client.Call("HelloService.Hello2", args, &reply)
			//if err != nil {
			//	log.Println("Client ERROR ", err)
			//}
			//log.Println("reply = ", reply)
			//log.Println(*calls)
		}()
		time.Sleep(time.Millisecond * 1000)
	}
}

func DemoRpc1() {

	go server()

	time.Sleep(time.Second * 1)

	//go clientGo()
	for i := 0; i < 1000; i++ {
		go clientCall()
	}
	//go client()
	//go client()

	select {}
}
