package demo1rpc

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (p *HelloService) Hello2(request map[string]string, reply *map[string]string) error {
	//log.Println("request = ", request)
	args := make(map[string]string)
	args["version"] = "8.2.0"
	args["user_id"] = "200"
	*reply = args
	//*reply = "hello:" + request
	time.Sleep(time.Second * 3) //
	return nil
}

func server() {

	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
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

func client() {
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
		log.Println("reply = ", reply)
		time.Sleep(time.Millisecond * 10)
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
			break
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

	go clientGo()
	//go client()
	//go client()
	//go client()

	select {}
}
