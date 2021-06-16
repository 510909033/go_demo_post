package main

import (
	"net/rpc"
	//"net/http"
	"log"
	"net"
	"time"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *([]string)) error {
	*reply = append(*reply, "test")
	time.Sleep(time.Second)
	return nil
}

func main() {
	newServer := rpc.NewServer()
	newServer.Register(new(Arith))

	l, e := net.Listen("tcp", "127.0.0.1:12341") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}

	go newServer.Accept(l)
	//	newServer.HandleHTTP("/foo", "/bar")
	time.Sleep(2 * time.Second)

	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:12341")
	if err != nil {
		panic(err)
	}
	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()
	ch := make(chan *rpc.Call, 3)
	for {
		args := &Args{7, 8}
		reply := make([]string, 10)
		//		err = client.Call("Arith.Multiply", args, &reply)

		client.Go("Arith.Multiply", args, &reply, ch)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		break
		//		log.Println(reply)
		//		time.Sleep(time.Millisecond)
	}
	msg := <-ch
	log.Println(msg.Reply)
}
