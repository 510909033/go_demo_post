package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	stProto "baotian0506.com/mygoprotobuf/proto"

	//protobuf编解码库,下面两个库是相互兼容的，可以使用其中任意一个
	"github.com/golang/protobuf/proto"
	//"github.com/gogo/protobuf/proto"
)

/*
开启一个 服务，
go run server_protobuf.go --ip localhost --port 12345
 */
func main() {
	ip := flag.String("ip", "localhost", "ip")
	port := flag.String("port", "6600", "端口")
	flag.Parse()

	log.Println("ip=", ip, *ip)
	log.Println("port=", port, *port)

	//监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *ip,*port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("new connect", conn.RemoteAddr())
		go readMessage(conn)
	}
}

//接收消息
func readMessage(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		//读消息
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Printf("断开连接？？， err=%+v", err)
			return
			//panic(err)
		}

		stReceive := &stProto.UserInfo{}
		pData := buf[:cnt]
		log.Printf("read原始内容=%s", string(pData))

		//protobuf解码
		err = proto.Unmarshal(pData, stReceive)
		if err != nil {
			panic(err)
		}

		log.Println("receive", conn.RemoteAddr(), stReceive)
		
		sendData,_:= proto.Marshal(&stProto.UserInfo{
			Cnt: 12345,
		})
		conn.Write(sendData)

		//conn.Write([]byte("hei哈dddddddddddddddddddddddddddddddddddddddddddddd"))
		if stReceive.Message == "stop" {
			os.Exit(1)
		}
	}
}
