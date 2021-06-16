package multi_groutinue

import (
	"context"
	"log"
	"runtime"
	"time"
)

func DemoMultiGroutinue() {
	multiSender_multi_Consume()
}

func multiSender_multi_Consume() {
	ch := make(chan int, 10)
	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	for i := 0; i < 10; i++ {
		go send(ch, ctx)
	}
	for i := 0; i < 10; i++ {
		go consume(ch, ctx)
	}

	cancel()

	go monitor(ctx)
	select {}
}

func closeCh(ch chan<- int, ctx context.Context) {

	close(ch)
}

func send(ch chan<- int, ctx context.Context) {
	defer func() {
		log.Println("send over")
	}()
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-timer.C:
			closeCh(ch, ctx)
			return
		case val := <-ticker.C:
			log.Println("before:", val.Second())
			ch <- val.Second()
			log.Println("af ter:", val.Second())
		}
	}
}

func consume(ch <-chan int, ctx context.Context) {
	defer func() {
		log.Println("consume over")
	}()
	for val := range ch {
		log.Println("consume:", val)
		time.Sleep(time.Second * 2)
	}

}

func monitor(ctx context.Context) {
	defer func() {
		log.Println("monitor over")
	}()
	for {
		log.Println(runtime.NumGoroutine())
		time.Sleep(time.Second * 3)
	}
}
