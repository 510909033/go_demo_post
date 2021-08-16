package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	stProto "baotian0506.com/mygoprotobuf/proto"

	//protobuf编解码库,下面两个库是相互兼容的，可以使用其中任意一个
	"github.com/golang/protobuf/proto"
	//"github.com/gogo/protobuf/proto"
)

func main() {

	var wg sync.WaitGroup
	size:= 10
	wg.Add(size)

	//并发连接到服务端
	for i:=0;i<size;i++{
		go func() {
			defer wg.Done()
			demo1()
		}()
	}

	log.Println("wg.waiting")
	wg.Wait()
	log.Println("main over")
}
func demo1() {
	strIP := "localhost:6600"
	var conn net.Conn
	var err error

	//连接服务器
	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	log.Println("connect", strIP, "success")
	defer conn.Close()

	//发送消息
	cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		stSend := &stProto.UserInfo{
			Message: sender.Text(),
			Length:  *proto.Int(len(sender.Text())),
			Cnt:     *proto.Int(cnt),
		}

		//protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		//发送
		writeLen, err := conn.Write(pData)
		log.Printf("writeLen , len=%v, err=%v",writeLen,err)


		var b = make([]byte, 20)
		n, err := conn.Read(b)
		log.Printf("red result, readLen=%v, err=%v",n,err)
		log.Printf("read content = %s ", string(b))
		if sender.Text() == "stop" {
			return
		}
	}
}
