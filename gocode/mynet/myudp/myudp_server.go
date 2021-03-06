package myudp

import (
	"log"
	"net"
)

//22
func DemoUdpServer() {
	// 监听UDP服务
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: PORT,
	})

	if err != nil {
		log.Fatal("Listen failed,", err)
		return
	}

	// 循环读取消息
	for {
		var data [1024]byte
		n, addr, err := udpConn.ReadFromUDP(data[:])
		if err != nil {
			log.Printf("Read from udp server:%s failed,err:%s", addr, err)
			break
		}
		go func() {
			// 返回数据
			log.Printf("Addr:%s,data:%v count:%d \n", addr, string(data[:n]), n)
			_, err := udpConn.WriteToUDP([]byte("msg recived."), addr)
			if err != nil {
				log.Println("write to udp server failed,err:", err)
			}
		}()
	}
	select {}
}
