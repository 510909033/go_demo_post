package my_groutinue

import (
	"log"
	"runtime"
	"time"
)

func DemoGroutinue() {
	multiSender_multi_Consume()
}

func multiSender_multi_Consume() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go send(ch)
	}
	for i := 0; i < 10; i++ {
		go consume(ch)
	}

	go monitor()
	select {}
}

func oneSender_multi_Consume() {
	ch := make(chan int, 10)
	go send(ch)

	for i := 0; i < 10; i++ {
		go consume(ch)
	}

	go monitor()
	select {}
}

func send(ch chan<- int) {
	defer func() {
		log.Println("send over")
	}()
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-timer.C:
			close(ch)
			return
		case val := <-ticker.C:
			log.Println("before:", val.Second())
			ch <- val.Second()
			log.Println("af ter:", val.Second())
		}
	}
}

func consume(ch <-chan int) {
	defer func() {
		log.Println("consume over")
	}()
	for val := range ch {
		log.Println("consume:", val)
		time.Sleep(time.Second * 2)
	}

}

func monitor() {
	defer func() {
		log.Println("monitor over")
	}()
	for {
		log.Println(runtime.NumGoroutine())
		time.Sleep(time.Second * 3)
	}
}
