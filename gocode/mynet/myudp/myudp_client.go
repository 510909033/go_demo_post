package myudp

import (
	"fmt"
	"log"
	"net"
	"time"
)

func DemoUdpClient() {
	// 连接服务器
	ts := time.Now()
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//IP:   net.IPv4(127, 0, 0, 1),
		IP:   net.IPv4(118, 125, 200, 123),
		Port: PORT,
	})
	log.Println("DialUDP 耗时：", time.Since(ts))

	if err != nil {
		log.Println("Connect to udp server failed,err:", err)
		return
	}

	for i := 0; i < 10; i++ {
		// 发送数据
		ts := time.Now()
		_, err := conn.Write([]byte(fmt.Sprintf("udp testing:%v", i)))
		log.Println("client Write 耗时：", time.Since(ts), "err=", err)
		if err != nil {
			log.Printf("Send data failed,err:", err)
			return
		}

		////接收数据
		//result := make([]byte, 1024)
		//ts = time.Now()
		//conn.SetReadDeadline(ts.Add(time.Second * 3))
		//n, remoteAddr, err := conn.ReadFromUDP(result)
		//log.Println("client ReadFromUDP 耗时：", time.Since(ts))
		//if err != nil {
		//	log.Printf("Read from udp server failed ,err:", err)
		//	return
		//}
		//log.Printf("Recived msg from %s, data:%s \n", remoteAddr, string(result[:n]))
	}
}
