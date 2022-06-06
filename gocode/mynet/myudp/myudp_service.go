package myudp

import (
	"log"
	"time"
)

const PORT = 9097

//21
func DemoUdpService() {

	//go DemoUdpServer()

	time.Sleep(time.Second)
	log.Println("server started...")

	go DemoUdpClient()

	time.Sleep(time.Second * 500)
	log.Println("DemoUdpService over")
}
