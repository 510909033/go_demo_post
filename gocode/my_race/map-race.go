package my_race

import (
	"log"
	"time"
)

func DemoNapace() {
	list := make(map[int]int)
	ch := make(chan struct{}, 20)

	i := 0
	for {
		i++
		if i > 10 {
			i = 0
		}
		ch <- struct{}{}
		go func(i int) {
			//i = 10
			defer func() {
				<-ch
			}()
			list[i] = i
			log.Printf("i=%d", i)
			time.Sleep(time.Second)
		}(i)

	}
}
